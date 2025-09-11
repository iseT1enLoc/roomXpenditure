package repository

import (
	"703room/703room.com/models"
	"context"
)

type RoomRepository interface {
	Create(ctx context.Context, room *models.Room) error
	Save(ctx context.Context, room *models.Room) error
	GetAllRooms(ctx context.Context) ([]models.Room, error)
	GetByID(ctx context.Context, id string) (*models.Room, error)
	ListByUserID(ctx context.Context, userID string) ([]models.Room, error)
	Update(ctx context.Context, room *models.Room) error
	Delete(ctx context.Context, id string) error
}
