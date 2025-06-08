package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Message struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    ChatID    primitive.ObjectID `bson:"chat_id" json:"chatId"`
    SenderID  int64              `bson:"sender_id" json:"senderId"`
    Content   string             `bson:"content" json:"content"`
    CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}
