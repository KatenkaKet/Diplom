package routes

import (
	"course-service/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	courseGroup := r.Group("/courses")
	{
		courseGroup.GET("/", controllers.GetCourses(db))
		courseGroup.GET("/:id", controllers.GetCourseByID(db))
		courseGroup.GET("/search", controllers.SearchCourses(db))
	}
}
