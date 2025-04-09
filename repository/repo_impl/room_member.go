package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"

	"gorm.io/gorm"
)

type roomMemberRepository struct {
	db *gorm.DB
}

func NewRoomMemberRepository(db *gorm.DB) repository.RoomMemberRepository {
	return &roomMemberRepository{db: db}
}

// AddMember inserts a new RoomMember
func (r *roomMemberRepository) AddMember(ctx context.Context, member *models.RoomMember) error {
	return r.db.WithContext(ctx).Create(member).Error
}

// GetMember retrieves a RoomMember by roomID and userID
func (r *roomMemberRepository) GetMember(ctx context.Context, roomID, userID string) (*models.RoomMember, error) {
	var member models.RoomMember
	if err := r.db.WithContext(ctx).
		Where("room_id = ? AND user_id = ?", roomID, userID).
		First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

// ListMembersByRoomID returns all members for a given room
func (r *roomMemberRepository) ListMembersByRoomID(ctx context.Context, roomID string) ([]models.RoomMember, error) {
	var members []models.RoomMember
	if err := r.db.WithContext(ctx).
		Where("room_id = ?", roomID).
		Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

// RemoveMember deletes a RoomMember by roomID and userID
func (r *roomMemberRepository) RemoveMember(ctx context.Context, roomID, userID string) error {
	return r.db.WithContext(ctx).
		Where("room_id = ? AND user_id = ?", roomID, userID).
		Delete(&models.RoomMember{}).Error
}

// UpdateRole updates the role of a RoomMember
func (r *roomMemberRepository) UpdateRole(ctx context.Context, roomID, userID, role string) error {
	return r.db.WithContext(ctx).
		Model(&models.RoomMember{}).
		Where("room_id = ? AND user_id = ?", roomID, userID).
		Update("role", role).Error
}
