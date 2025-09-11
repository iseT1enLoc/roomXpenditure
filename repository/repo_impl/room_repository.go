package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"
	"log"

	"gorm.io/gorm"
)

type roomRepository struct {
	db *gorm.DB
}

// NewRoomRepository returns a new RoomRepository instance
func NewRoomRepository(db *gorm.DB) repository.RoomRepository {
	return &roomRepository{db: db}
}

// Create inserts a new room into the database
func (r *roomRepository) Create(ctx context.Context, room *models.Room) error {
	return r.db.WithContext(ctx).Create(room).Error
}

// GetByID retrieves a room by its ID
func (r *roomRepository) GetByID(ctx context.Context, id string) (*models.Room, error) {
	var room models.Room
	log.Println("Enter line 29 of get by id")
	err := r.db.WithContext(ctx).
		Preload("Members", "room_id = ?", id).
		Find(&room)
	if err != nil {
		return nil, err.Error
	}
	return &room, nil
}

func (r *roomRepository) ListByUserID(ctx context.Context, userID string) ([]models.Room, error) {
	var rooms []models.Room
	err := r.db.WithContext(ctx).
		Model(&models.Room{}).
		// get all rooms where this user is a member
		Joins("JOIN room_members rm ON rm.room_id = rooms.room_id").
		Where("rm.user_id = ?", userID).
		// preload all members of the room
		Preload("Members").
		Find(&rooms).Error

	log.Println(err)
	return rooms, nil
}

// Update updates room details
func (r *roomRepository) Update(ctx context.Context, room *models.Room) error {
	return r.db.WithContext(ctx).Save(room).Error
}

// Delete removes a room by ID
func (r *roomRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.Room{}, "room_id = ?", id).Error
}
