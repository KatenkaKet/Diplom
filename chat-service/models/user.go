package models

type User struct {
    ID        int64  `bson:"id" json:"id"`
    Username  string `bson:"username" json:"username"`
    AvatarURL string `bson:"avatar_url" json:"avatar_url"`
}
