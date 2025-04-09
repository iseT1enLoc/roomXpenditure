package repository

import (
	"703room/703room.com/models"
	"context"
)

type RoomMemberRepository interface {
	AddMember(ctx context.Context, member *models.RoomMember) error
	GetMember(ctx context.Context, roomID, userID string) (*models.RoomMember, error)
	ListMembersByRoomID(ctx context.Context, roomID string) ([]models.RoomMember, error)
	RemoveMember(ctx context.Context, roomID, userID string) error
	UpdateRole(ctx context.Context, roomID, userID, role string) error
}
