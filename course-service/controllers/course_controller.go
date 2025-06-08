package controllers

import (
	"course-service/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCourses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var courses []models.Course
		if err := db.Preload("Chapters.Topics").Find(&courses).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch courses"})
			return
		}
		c.JSON(http.StatusOK, courses)
	}
}

func GetCourseByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
			return
		}

		var course models.Course
		if err := db.Preload("Chapters.Topics").First(&course, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			}
			return
		}

		c.JSON(http.StatusOK, course)
	}
}

func SearchCourses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
			return
		}

		var courses []models.Course
		searchQuery := "%" + strings.ToLower(query) + "%"

		if err := db.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", searchQuery, searchQuery).
			Preload("Chapters.Topics").
			Find(&courses).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search courses"})
			return
		}

		c.JSON(http.StatusOK, courses)
	}
}
