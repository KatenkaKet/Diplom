package routes

import (
    "auth-service/controllers"
    "auth-service/middlewares"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    api := router.Group("/api")
    {
        api.POST("/register", func(c *gin.Context) {
            controllers.Register(c, db)
        })

        api.POST("/login", func(c *gin.Context) {
            controllers.Login(c, db)
        })

        // Защищённые маршруты (требуют JWT)
        auth := api.Group("/")
        auth.Use(middlewares.JWTAuthMiddleware())
        {
            auth.GET("/profile", func(c *gin.Context) {
                controllers.GetProfile(c, db)
            })
        }
    }

    RegisterUserRoutes(router, db)
}

