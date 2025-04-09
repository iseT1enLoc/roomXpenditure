package routes

import (
	"703room/703room.com/api/handlers"
	repoimpl "703room/703room.com/repository/repo_impl"
	serviceimpl "703room/703room.com/services/service_impl"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoute(timeout time.Duration, db *gorm.DB, r *gin.RouterGroup) {
	user_repo := repoimpl.NewUserRepository(db)
	user_service := serviceimpl.NewUserService(user_repo)
	auth_service := serviceimpl.NewAuthService(user_repo)
	handler := handlers.NewUserHandler(auth_service, user_service)
	r.POST("/login", handler.Login)
	r.POST("/signup", handler.Signup())
	r.GET("/user/me", handler.GetCurrentUser)
}
