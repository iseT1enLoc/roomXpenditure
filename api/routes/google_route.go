package routes

import (
	"os"
	"time"

	"703room/703room.com/api/handlers"
	repoimpl "703room/703room.com/repository/repo_impl"
	serviceimpl "703room/703room.com/services/service_impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewGoogleRouter(timeout time.Duration, db *gorm.DB, r *gin.RouterGroup, p *gin.RouterGroup) {
	user_repo := repoimpl.NewUserRepository(db)
	auth_service := serviceimpl.NewAuthService(user_repo)
	google_service := serviceimpl.NewGoogleService(user_repo, auth_service)
	client_id := os.Getenv("GOOGLE_CLIENT_ID")
	client_secret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirect_url := os.Getenv("REDIRECT_URL")
	handler := handlers.NewGoogleHandler(google_service, client_id, client_secret, redirect_url)
	r.GET("/google/login", handler.RedirectToGoogle)
	r.GET("/google/callback", handler.HandleGoogleCallback)
}
