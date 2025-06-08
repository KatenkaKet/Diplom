package controllers

import (
    // "auth-service/database"
    "auth-service/models"
    "net/http"
    "strconv"
	"gorm.io/gorm"

    "github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context, db *gorm.DB) {
    idParam := c.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }

    var user models.User
    if err := db.Where("id = ?", id).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":         user.ID,
        "username":   user.Username,
        "avatar_url": user.AvatarURL,
    })
}
