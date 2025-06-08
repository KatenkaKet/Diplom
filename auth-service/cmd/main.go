package main

import (
    "auth-service/config"
    "auth-service/database"
    "auth-service/routes"
    "github.com/gin-gonic/gin"
    "log"
	"github.com/gin-contrib/cors"
)

func main() {
    // Загружаем переменные из .env
    config.LoadEnv()

    // Подключаемся к базе данных
    db := database.ConnectDB()
    
    // Создаем роутер
    router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))


    // Настраиваем маршруты
    routes.SetupRoutes(router, db)

    // Запуск сервера
    port := config.GetEnv("PORT", "8080")
    log.Println("🚀 Server is running on port " + port)
    router.Run(":" + port)
}
