package models

import (
    "time"
)

type Course struct {
    ID               uint      `gorm:"primaryKey"`
    Title            string    `gorm:"not null"`
    ShortDescription string
    Description      string
    Outcomes         string
    Audience         string
    AboutAuthor      string
    Category         string
    Price            float64
    ImageURL         string
    AuthorID         uint      //`gorm:"not null"`
    CreatedAt        time.Time
    UpdatedAt        time.Time
    Chapters         []Chapter `gorm:"foreignKey:CourseID"`
}
