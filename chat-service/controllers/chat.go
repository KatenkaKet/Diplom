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

func GetChats(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID type"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"members": userIDInt64}
	opts := options.Find().SetSort(bson.M{"updated_at": -1})

	cursor, err := database.ChatCollection.Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch chats"})
		return
	}

	var chats []models.Chat
	if err = cursor.All(ctx, &chats); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode chats"})
		return
	}

	// Если чатов нет, возвращаем пустой массив
	if len(chats) == 0 {
		c.JSON(http.StatusOK, []interface{}{})
		return
	}

	type PartnerInfo struct {
		ID        int64  `json:"id"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	}

	type ChatWithPartner struct {
		ID          primitive.ObjectID `json:"id"`
		Type        string             `json:"type"`
		LastMessage string             `json:"last_message"`
		UpdatedAt   time.Time          `json:"updated_at"`
		Partner     PartnerInfo        `json:"partner"`
	}

	var result []ChatWithPartner
	for _, chat := range chats {
		var partnerID int64
		for _, id := range chat.Members {
			if id != userIDInt64 {
				partnerID = id
				break
			}
		}

		partner, err := external.GetUserByID2(partnerID)
		if err != nil {
			// Пропускаем чат, если не удалось получить информацию о партнере
			continue
		}

		result = append(result, ChatWithPartner{
			ID:          chat.ID,
			Type:        string(chat.Type),
			LastMessage: chat.LastMessage,
			UpdatedAt:   chat.UpdatedAt,
			Partner: PartnerInfo{
				ID:        partner.ID,
				Username:  partner.Username,
				AvatarURL: partner.AvatarURL,
			},
		})
	}

	c.JSON(http.StatusOK, result)
}

type CreateChatInput struct {
	Type      models.ChatType `json:"type" binding:"required"`
	Title     string          `json:"title"`
	Members   []int64         `json:"members" binding:"required"`
	Admins    []int64         `json:"admins"`
	AvatarURL string          `json:"avatar_url"`
}

func CreateChat(c *gin.Context) {
	var input CreateChatInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Type == models.ChatTypePrivate && len(input.Members) == 2 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		filter := bson.M{
			"type":    models.ChatTypePrivate,
			"members": bson.M{"$all": input.Members, "$size": 2},
		}

		var existing models.Chat
		err := database.ChatCollection.FindOne(ctx, filter).Decode(&existing)
		if err == nil {
			c.JSON(http.StatusOK, existing)
			return
		}
	}

	chat := models.Chat{
		ID:          primitive.NewObjectID(),
		Type:        input.Type,
		Title:       input.Title,
		AvatarURL:   input.AvatarURL,
		Members:     input.Members,
		Admins:      input.Admins,
		LastMessage: "",
		UpdatedAt:   time.Now(),
	}

	_, err := database.ChatCollection.InsertOne(context.Background(), chat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create chat"})
		return
	}

	c.JSON(http.StatusCreated, chat)
}
