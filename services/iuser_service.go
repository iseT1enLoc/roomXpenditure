package services

import (
	"703room/703room.com/models"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	// User core
	RegisterUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUsersByRoomID(ctx context.Context, roomID uuid.UUID) ([]models.User, error)
	GetAllUserRoomsByUserID(ctx context.Context, roomID uuid.UUID) ([]models.Room, error)
}
