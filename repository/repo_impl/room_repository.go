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
	if err := r.db.WithContext(ctx).Preload("RoomMembers").Preload("Expenses").First(&room, "room_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}

// ListByUserID lists rooms where the user is a member
// ListByUserID lists rooms where the user is a member
func (r *roomRepository) ListByUserID(ctx context.Context, userID string) ([]models.Room, error) {
	var rooms []models.Room

	// Log the input
	log.Println("ListByUserID called with userID:", userID)

	tx := r.db.WithContext(ctx).
		Table("rooms").
		Joins("JOIN room_members ON room_members.room_id = rooms.room_id").
		Where("room_members.user_id = ?", userID)

	// Log the generated SQL (if needed)
	stmt := tx.Statement
	tx.Find(&rooms)

	log.Println("Executed SQL:", stmt.SQL.String(), "with Vars:", stmt.Vars)
	log.Println("Rooms found:", len(rooms))

	if tx.Error != nil {
		log.Println("Error while fetching rooms:", tx.Error)
		return nil, tx.Error
	}

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
