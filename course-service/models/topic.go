package models

type Topic struct {
    ID        uint   `gorm:"primaryKey"`
    ChapterID uint   `gorm:"not null"`
    Title     string `gorm:"not null"`
    Content   string
    Position  int    `gorm:"default:0"`
}
