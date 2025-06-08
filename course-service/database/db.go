package database

import (
    "course-service/config"
    "course-service/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func InitDB(cfg config.Config) *gorm.DB {
    db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }

    err = db.AutoMigrate(&models.Course{}, &models.Chapter{}, &models.Topic{})
    if err != nil {
        log.Fatal("failed to migrate database schema: ", err)
    }

    return db
}