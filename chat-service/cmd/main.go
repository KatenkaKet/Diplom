package main

import (
	"chat-service/database"
	"chat-service/external"
	"chat-service/migrations"
	"chat-service/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env файл не найден, используются переменные окружения по умолчанию")
	}
}


func main() {
	// Подключение к MongoDB
	database.ConnectMongo()

	// Подключение к PostgreSQL
	if err := external.ConnectAuthDB(); err != nil {
		log.Fatalf("❌ Ошибка подключения к auth-db: %v", err)
	}

	// Миграции MongoDB
	migrations.Run()

	// Инициализация сервера
	r := gin.Default()

	// ✅ CORS должен идти до SetupRoutes
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Маршруты
	routes.SetupRoutes(r)

	// Старт
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Printf("🚀 Сервер запущен на http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
