package repository

import (
	"703room/703room.com/models"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	GetUsersByRoomID(ctx context.Context, roomID uuid.UUID) ([]models.User, error)
	GetAllUserRoomsByUserID(ctx context.Context, roomID uuid.UUID) ([]models.Room, error)
}
