package routes

import (
	"time"

	"703room/703room.com/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUp(timeout time.Duration, db *gorm.DB, r *gin.Engine) {
	publicRoute := r.Group("/api/public")       //ai vo cung duoc
	protectedRoute := r.Group("/api/protected") // co account moi vo duoc
	protectedRoute.Use(middlewares.JWTMiddleware())

	NewUserRoute(100*time.Second, db, publicRoute, protectedRoute)

}
