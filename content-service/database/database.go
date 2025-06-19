// content-service/database/database.go
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

var (
    Client            *mongo.Client
    ChapterCollection *mongo.Collection
    TopicCollection   *mongo.Collection
    ContentCollection *mongo.Collection
    CommentCollection *mongo.Collection
)

// ConnectMongo читает MONGO_URI и MONGO_DB_NAME из окружения,
// подключается к MongoDB и инициализирует коллекции.
func ConnectMongo() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // URI для подключения (или localhost по умолчанию)
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017"
    }

    // Имя базы (или content-service по умолчанию)
    dbName := os.Getenv("MONGO_DB_NAME")
    if dbName == "" {
        dbName = "content-service"
    }

    clientOpts := options.Client().ApplyURI(mongoURI)
    var err error
    Client, err = mongo.Connect(ctx, clientOpts)
    if err != nil {
        log.Fatalf("MongoDB connection error: %v", err)
    }

    if err = Client.Ping(ctx, nil); err != nil {
        log.Fatalf("MongoDB ping failed: %v", err)
    }

    fmt.Printf("✅ Connected to MongoDB at %s, using database %s\n", mongoURI, dbName)

    db := Client.Database(dbName)
    ChapterCollection = db.Collection("chapters")
    TopicCollection = db.Collection("topics")
    ContentCollection = db.Collection("content")
    CommentCollection = db.Collection("comments")
}
