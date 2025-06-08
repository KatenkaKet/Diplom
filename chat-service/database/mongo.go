package database

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var ChatCollection *mongo.Collection
var MessageCollection *mongo.Collection

func ConnectMongo() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017"
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    var err error
    Client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Mongo connection error: %v", err)
    }

    // Проверим соединение
    if err := Client.Ping(ctx, nil); err != nil {
        log.Fatalf("Mongo ping failed: %v", err)
    }

    fmt.Println("✅ Connected to MongoDB")

    db := Client.Database("chat-service")
    ChatCollection = db.Collection("chats")
    MessageCollection = db.Collection("messages")
}
