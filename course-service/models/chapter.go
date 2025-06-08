package models

type Chapter struct {
    ID       uint     `gorm:"primaryKey"`
    CourseID uint     `gorm:"not null"`
    Title    string   `gorm:"not null"`
    Position int      `gorm:"default:0"`
    Topics   []Topic  `gorm:"foreignKey:ChapterID"`
}
