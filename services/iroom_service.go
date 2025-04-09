package services

import (
	"703room/703room.com/models"
	"context"
)

type RoomService interface {
	CreateRoom(ctx context.Context, room *models.Room) error
	GetRoomByID(ctx context.Context, id string) (*models.Room, error)
	ListRoomsByUserID(ctx context.Context, userID string) ([]models.Room, error)
	UpdateRoom(ctx context.Context, room *models.Room) error
	DeleteRoom(ctx context.Context, id string) error
	AddMember(ctx context.Context, member *models.RoomMember) error
	GetMember(ctx context.Context, roomID, userID string) (*models.RoomMember, error)
	ListMembersByRoomID(ctx context.Context, roomID string) ([]models.RoomMember, error)
	RemoveMember(ctx context.Context, roomID, userID string) error
	UpdateRole(ctx context.Context, roomID, userID, role string) error
}
