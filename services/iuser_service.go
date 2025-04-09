package services

import (
	"703room/703room.com/models"
	"context"
)

type UserService interface {
	// User core
	RegisterUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error

	// Room manipulation
	CreateRoom(ctx context.Context, userID string, roomName string) (*models.Room, error)
	JoinRoom(ctx context.Context, userID, roomID string) error
	LeaveRoom(ctx context.Context, userID, roomID string) error
	ListUserRooms(ctx context.Context, userID string) ([]models.Room, error)

	//expense
	AddExpense(ctx context.Context, expense *models.Expense) error
	UpdateExpense(ctx context.Context, expense *models.Expense) error
	DeleteExpense(ctx context.Context, id string) error
	GetExpensesByRoomID(ctx context.Context, roomID string) ([]models.Expense, error)
}
