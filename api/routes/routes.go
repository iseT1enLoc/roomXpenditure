package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUp(timeout time.Duration, db *gorm.DB, r *gin.Engine) {
	publicRoute := r.Group("/api/public") //ai vo cung duoc
	//protectedRoute := r.Group("/api/protected") // co account moi vo duoc

	NewUserRoute(100*time.Second, db, publicRoute)
}
