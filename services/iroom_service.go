package services

import (
	"703room/703room.com/models"
	"context"

	"github.com/google/uuid"
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
	SendInvitationToUsers(ctx context.Context, fromUserID, roomID uuid.UUID, emails []string, message string) error
	GetAllPendingInvitationByUserId(ctx context.Context, user_id uuid.UUID) ([]models.RoomExpenseInvitationRecipient, error)
	UpdateInvitationRequestStatus(ctx context.Context, recipientID uuid.UUID, status models.InvitationStatus) error
	UpdateMemberCount(ctx context.Context) error
}
