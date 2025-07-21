package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(gin *gin.Engine, gormDB *gorm.DB) {
	// @BasePath /api/v1
	Router := gin.Group("/api/v1")
	NewBookRouter(Router, gormDB)
}
