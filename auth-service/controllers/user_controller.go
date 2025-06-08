package controllers

import (
    "auth-service/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

func GetProfile(c *gin.Context, db *gorm.DB) {
    userID := c.MustGet("user_id").(uint)

    var user models.User
    if err := db.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":         user.ID,
        "first_name": user.FirstName,
        "last_name":  user.LastName,
        "email":      user.Email,
        "username":   user.Username,
        "phone":      user.Phone,
        "avatar_url": user.AvatarURL,
    })
}
