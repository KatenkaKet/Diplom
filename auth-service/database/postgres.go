package database

import (
    "auth-service/config"
    "auth-service/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func ConnectDB() *gorm.DB {
    dsn := "host=" + config.GetEnv("DB_HOST", "localhost") +
        " user=" + config.GetEnv("DB_USER", "postgres") +
        " password=" + config.GetEnv("DB_PASSWORD", "password") +
        " dbname=" + config.GetEnv("DB_NAME", "auth_service") +
        " port=" + config.GetEnv("DB_PORT", "5432") +
        " sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Не удалось подключиться к базе данных:", err)
    }

    db.AutoMigrate(&models.User{})
    return db
}
