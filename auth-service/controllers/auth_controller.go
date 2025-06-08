package controllers

import (
	"auth-service/models"
	"auth-service/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func Register(c *gin.Context, db *gorm.DB) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования пароля"})
		return
	}

	user := models.User{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		MiddleName: input.MiddleName,
		Email:      input.Email,
		Phone:      input.Phone,
		Username:   input.Username,
		Password:   hashedPassword, // Сохраняем хеш
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context, db *gorm.DB) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}
	token, err := utils.GenerateJWT(user.ID, os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
