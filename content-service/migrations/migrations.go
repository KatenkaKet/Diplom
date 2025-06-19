package migrations

import (
    "context"
    "fmt"
    "log"
    "time"

    "content-service/database"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Run применяет все индексы для коллекций chapters, topics и comments
func Run() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // 1) Chapters: индекс по id_course
    chapterIdx := mongo.IndexModel{
        Keys:    bson.D{{Key: "id_course", Value: 1}},
        Options: options.Index().SetName("idx_chapters_id_course").SetBackground(true),
    }
    if _, err := database.ChapterCollection.Indexes().CreateOne(ctx, chapterIdx); err != nil {
        log.Fatalf("Не удалось создать индекс для chapters: %v", err)
    }

    // 2) Topics: 
    //    a) индекс по id_course
    //    b) индекс по id_chapter
    //    c) составной индекс по storyline.type (для быстрых выборок по типу контента)
    topicIndexes := []mongo.IndexModel{
        {
            Keys:    bson.D{{Key: "id_course", Value: 1}},
            Options: options.Index().SetName("idx_topics_id_course").SetBackground(true),
        },
        {
            Keys:    bson.D{{Key: "id_chapter", Value: 1}},
            Options: options.Index().SetName("idx_topics_id_chapter").SetBackground(true),
        },
        {
            Keys:    bson.D{{Key: "storyline.type", Value: 1}},
            Options: options.Index().SetName("idx_topics_storyline_type").SetBackground(true),
        },
    }
    if _, err := database.TopicCollection.Indexes().CreateMany(ctx, topicIndexes); err != nil {
        log.Fatalf("Не удалось создать индексы для topics: %v", err)
    }

    // 3) Comments:
    //    a) индекс по id_topic
    //    b) индекс по id_user
    //    c) индекс по parent_id (для быстрого строения древовидных комментариев)
    commentIndexes := []mongo.IndexModel{
        {
            Keys:    bson.D{{Key: "id_topic", Value: 1}},
            Options: options.Index().SetName("idx_comments_id_topic").SetBackground(true),
        },
        {
            Keys:    bson.D{{Key: "id_user", Value: 1}},
            Options: options.Index().SetName("idx_comments_id_user").SetBackground(true),
        },
        {
            Keys:    bson.D{{Key: "parent_id", Value: 1}},
            Options: options.Index().SetName("idx_comments_parent_id").SetBackground(true),
        },
    }
    if _, err := database.CommentCollection.Indexes().CreateMany(ctx, commentIndexes); err != nil {
        log.Fatalf("Не удалось создать индексы для comments: %v", err)
    }

    fmt.Println("✅ Mongo-миграции content-service успешно применены")
}
