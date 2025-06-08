package routes

import (
    "auth-service/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
    r.GET("/api/users/:id", func(c *gin.Context) {
        controllers.GetUserByID(c, db)
    })
}
