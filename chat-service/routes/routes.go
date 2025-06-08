package routes

import (
    "chat-service/controllers"
    "github.com/gin-gonic/gin"
	"chat-service/middleware"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    protected := r.Group("/api", middleware.JWTMiddleware())
	
    protected.POST("/messages", controllers.SendMessage)
    protected.POST("/chats", controllers.CreateChat)
    protected.GET("/chats", controllers.GetChats)
    protected.GET("/chats/:id/messages", controllers.GetMessagesByChat)

    protected.GET("/users/search", controllers.SearchUsers)

}
