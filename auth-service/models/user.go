package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    FirstName  string
    LastName   string
    MiddleName string
    Email      string `gorm:"unique;not null"`
    Phone      string
    Username   string `gorm:"unique;not null"`
    Password   string `gorm:"not null"`
    AvatarURL  string
}
