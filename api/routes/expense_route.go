package routes

import (
	"703room/703room.com/api/handlers"
	repoimpl "703room/703room.com/repository/repo_impl"
	serviceimpl "703room/703room.com/services/service_impl"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewExpenseRoute(timeout time.Duration, db *gorm.DB, r *gin.RouterGroup, p *gin.RouterGroup) {
	expense_repo := repoimpl.NewExpenseRepository(db)
	expense_service := serviceimpl.NewExpenseService(expense_repo)
	user_repo := repoimpl.NewUserRepository(db)
	user_service := serviceimpl.NewUserService(user_repo)
	handler := handlers.NewExpenseHandler(expense_service, user_service)
	//r.Use(middlewares.JWTMiddleware(auth_service))
	p.POST("/expense", handler.CreateNewExpense())
	p.GET("/expense/:id", handler.GetExpenseByID())
	p.GET("/expense/user", handler.GetExpensesFiltered())
	p.GET("/expense/calc", handler.CalculateMonthExpense())
	// p.GET("/user/me", handler.GetCurrentUser())
}
