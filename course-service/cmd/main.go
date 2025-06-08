package main

import (
	"course-service/config"
	"course-service/database"
	"course-service/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Загрузка конфигурации из .env
	cfg := config.LoadConfig()

	// Подключение к базе данных и миграции
	db := database.InitDB(cfg)

	// Инициализация роутера
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Регистрация маршрутов
	api := r.Group("/api")
	routes.RegisterRoutes(api, db)

	// Проверка состояния сервиса
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "course-service is running"})
	})

	// Запуск сервера
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
