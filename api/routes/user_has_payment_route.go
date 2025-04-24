package routes

import (
	"703room/703room.com/api/handlers"
	repoimpl "703room/703room.com/repository/repo_impl"
	serviceimpl "703room/703room.com/services/service_impl"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserHasPaymentRoute(timeout time.Duration, db *gorm.DB, r *gin.RouterGroup, p *gin.RouterGroup) {
	user_has_payment_repo := repoimpl.NewUserHasPaymentRepository(db)
	user_has_payment_service := serviceimpl.NewUserHasPaymentService(user_has_payment_repo)
	user_repo := repoimpl.NewUserRepository(db)
	user_service := serviceimpl.NewUserService(user_repo)
	handler := handlers.NewUserHasPaymentHandler(user_has_payment_service, user_service)

	//r.Use(middlewares.JWTMiddleware(auth_service))
	p.POST("/expense", handler.CreateNewExpense())
	//p.GET("/expense/:id", handler.GetExpenseByID())
	p.GET("/expense/user", handler.GetExpensesFiltered())
	log.Println("Enter user has payment route")
	p.GET("/expense/calc", handler.CalculateMonthExpense())
	// p.GET("/user/me", handler.GetCurrentUser())
}
