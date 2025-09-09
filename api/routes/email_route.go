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

func NewEmailRoute(timeout time.Duration, db *gorm.DB, r *gin.RouterGroup, p *gin.RouterGroup) {
	user_repo := repoimpl.NewUserRepository(db)
	expense_repo := repoimpl.NewUserHasPaymentRepository(db)
	room_repo := repoimpl.NewRoomRepository(db)
	room_member_repo := repoimpl.NewRoomMemberRepository(db)

	invitation_repo := repoimpl.NewInvitationRepo(db)

	roomService := serviceimpl.NewRoomService(room_repo, room_member_repo, user_repo, invitation_repo)
	expenseService := serviceimpl.NewUserHasPaymentService(expense_repo)
	userService := serviceimpl.NewUserService(user_repo)

	EMAIL_API := os.Getenv("EMAIL_API")
	email_service := serviceimpl.NewEmailService(EMAIL_API, expenseService, roomService, userService)
	handler := handlers.NewEmailhandler(email_service)
	r.POST("/send-email", handler.SendEmailHandler())
}
