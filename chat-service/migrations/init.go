package migrations

import (
    "chat-service/database"
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
    "log"
    "time"
)

func Run() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    index := mongoIndex("members_index", "members")
    _, err := database.ChatCollection.Indexes().CreateOne(ctx, index)
    if err != nil {
        log.Fatalf("Ошибка создания индекса: %v", err)
    }

    fmt.Println("✅ Миграции успешно применены")
}

func mongoIndex(name, field string) mongo.IndexModel {
    return mongo.IndexModel{
        Keys:    bson.D{{Key: field, Value: 1}},
        Options: options.Index().SetName(name).SetBackground(true),
    }
}
