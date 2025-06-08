package controllers

import (
    "chat-service/database"
    "chat-service/external"
    "chat-service/models"
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MessageInput struct {
    ChatID  string `json:"chat_id" binding:"required"`
    Content string `json:"content" binding:"required"`
}

type MessageResponse struct {
    ID        interface{} `json:"id"`
    ChatID    interface{} `json:"chat_id"`
    SenderID  int64       `json:"sender_id"`
    Content   string      `json:"content"`
    CreatedAt time.Time   `json:"created_at"`
    Sender    struct {
        Username  string `json:"username"`
        AvatarURL string `json:"avatar_url"`
    } `json:"sender"`
}

func SendMessage(c *gin.Context) {
    var input MessageInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    chatID, err := primitive.ObjectIDFromHex(input.ChatID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chat ID"})
        return
    }

    senderID := c.GetInt64("user_id")

    // Проверка: участвует ли пользователь в чате
    var chat models.Chat
    err = database.ChatCollection.FindOne(context.Background(), bson.M{"_id": chatID}).Decode(&chat)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "chat not found"})
        return
    }

    allowed := false
    for _, memberID := range chat.Members {
        if memberID == senderID {
            allowed = true
            break
        }
    }

    if !allowed {
        c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this chat"})
        return
    }

    user, err := external.GetUserByID(senderID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "sender not found"})
        return
    }

    msg := models.Message{
        ChatID:    chatID,
        SenderID:  senderID,
        Content:   input.Content,
        CreatedAt: time.Now(),
    }

    res, err := database.MessageCollection.InsertOne(context.Background(), msg)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save message"})
        return
    }

    // Обновить чат
    database.ChatCollection.UpdateByID(context.Background(), chatID, bson.M{
        "$set": bson.M{
            "last_message": input.Content,
            "updated_at":   time.Now(),
        },
    })

    c.JSON(http.StatusOK, gin.H{
        "id":         res.InsertedID,
        "chat_id":    input.ChatID,
        "sender_id":  senderID,
        "content":    input.Content,
        "created_at": msg.CreatedAt,
        "sender": gin.H{
            "username":   user.Username,
            "avatar_url": user.AvatarURL,
        },
    })
}


func GetMessagesByChat(c *gin.Context) {
    chatIDHex := c.Param("id")
    chatID, err := primitive.ObjectIDFromHex(chatIDHex)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chat ID"})
        return
    }

    userID := c.GetInt64("user_id")

    // Найдём чат и проверим, что userID входит в members
    var chat models.Chat
    err = database.ChatCollection.FindOne(context.Background(), bson.M{"_id": chatID}).Decode(&chat)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "chat not found"})
        return
    }

    allowed := false
    for _, id := range chat.Members {
        if id == userID {
            allowed = true
            break
        }
    }

    if !allowed {
        c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
        return
    }

    // Загружаем сообщения
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"chat_id": chatID}
    opts := options.Find().SetSort(bson.M{"created_at": 1})

    cursor, err := database.MessageCollection.Find(ctx, filter, opts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
        return
    }

    var messages []models.Message
    if err := cursor.All(ctx, &messages); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "decode error"})
        return
    }

    var result []MessageResponse
    for _, msg := range messages {
        user, _ := external.GetUserByID(msg.SenderID)

        r := MessageResponse{
            ID:        msg.ID,
            ChatID:    msg.ChatID,
            SenderID:  msg.SenderID,
            Content:   msg.Content,
            CreatedAt: msg.CreatedAt,
        }

        if user != nil {
            r.Sender.Username = user.Username
            r.Sender.AvatarURL = user.AvatarURL
        }

        result = append(result, r)
    }

    c.JSON(http.StatusOK, result)
}
