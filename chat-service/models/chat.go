package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type ChatType string

const (
    ChatTypePrivate ChatType = "private"
    ChatTypeGroup   ChatType = "group"
    ChatTypeChannel ChatType = "channel"
)

type Chat struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Type        ChatType           `bson:"type" json:"type"`
    Title       string             `bson:"title,omitempty" json:"title"`
    AvatarURL   string             `bson:"avatar_url,omitempty" json:"avatarUrl"`
    Members     []int64            `bson:"members" json:"members"`
    Admins      []int64            `bson:"admins,omitempty" json:"admins"`
    LastMessage string             `bson:"last_message,omitempty" json:"lastMessage"`
    UpdatedAt   time.Time          `bson:"updated_at" json:"updatedAt"`
}
