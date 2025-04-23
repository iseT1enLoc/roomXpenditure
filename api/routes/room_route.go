package routes

import (
	"703room/703room.com/api/handlers"
	repoimpl "703room/703room.com/repository/repo_impl"
	serviceimpl "703room/703room.com/services/service_impl"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoomRoute(timeout time.Duration, db *gorm.DB, r *gin.RouterGroup, p *gin.RouterGroup) {
	room_repo := repoimpl.NewRoomRepository(db)
	room_member_repo := repoimpl.NewRoomMemberRepository(db)
	room_service := serviceimpl.NewRoomService(room_repo, room_member_repo)
	handler := handlers.NewRoomHandler(room_service)

	p.POST("/room/create", handler.CreateNewRoom())
}
