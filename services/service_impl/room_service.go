package serviceimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type roomService struct {
	roomRepo       repository.RoomRepository
	roomMemberRepo repository.RoomMemberRepository
}

func NewRoomService(roomRepo repository.RoomRepository, roomMemberRepo repository.RoomMemberRepository) services.RoomService {
	return &roomService{
		roomRepo:       roomRepo,
		roomMemberRepo: roomMemberRepo,
	}
}

// Room CRUD

func (s *roomService) CreateRoom(ctx context.Context, room *models.Room) error {
	if room == nil {
		return errors.New("room cannot be nil")
	}
	if strings.TrimSpace(room.RoomName) == "" {
		return errors.New("room name cannot be empty")
	}
	if room.CreatedBy == uuid.Nil {
		return errors.New("created_by must be a valid UUID")
	}
	return s.roomRepo.Create(ctx, room)
}

func (s *roomService) GetRoomByID(ctx context.Context, id string) (*models.Room, error) {
	if id == "" {
		return nil, errors.New("room ID is required")
	}
	return s.roomRepo.GetByID(ctx, id)
}

func (s *roomService) ListRoomsByUserID(ctx context.Context, userID string) ([]models.Room, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	return s.roomRepo.ListByUserID(ctx, userID)
}

func (s *roomService) UpdateRoom(ctx context.Context, room *models.Room) error {
	if room == nil || room.RoomID == uuid.Nil {
		return errors.New("invalid room data")
	}
	return s.roomRepo.Update(ctx, room)
}

func (s *roomService) DeleteRoom(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("room ID is required")
	}
	return s.roomRepo.Delete(ctx, id)
}

// Room Members

func (s *roomService) AddMember(ctx context.Context, member *models.RoomMember) error {
	if member == nil {
		return errors.New("member cannot be nil")
	}
	if member.RoomID == uuid.Nil || member.UserID == uuid.Nil {
		return errors.New("room ID and user ID must be valid UUIDs")
	}
	if member.Role == "" {
		member.Role = "member"
	}
	return s.roomMemberRepo.AddMember(ctx, member)
}

func (s *roomService) GetMember(ctx context.Context, roomID, userID string) (*models.RoomMember, error) {
	if roomID == "" || userID == "" {
		return nil, errors.New("room ID and user ID are required")
	}
	return s.roomMemberRepo.GetMember(ctx, roomID, userID)
}

func (s *roomService) ListMembersByRoomID(ctx context.Context, roomID string) ([]models.RoomMember, error) {
	if roomID == "" {
		return nil, errors.New("room ID is required")
	}
	return s.roomMemberRepo.ListMembersByRoomID(ctx, roomID)
}

func (s *roomService) RemoveMember(ctx context.Context, roomID, userID string) error {
	if roomID == "" || userID == "" {
		return errors.New("room ID and user ID are required")
	}
	return s.roomMemberRepo.RemoveMember(ctx, roomID, userID)
}

func (s *roomService) UpdateRole(ctx context.Context, roomID, userID, role string) error {
	if roomID == "" || userID == "" {
		return errors.New("room ID and user ID are required")
	}
	if role == "" {
		return errors.New("role is required")
	}
	return s.roomMemberRepo.UpdateRole(ctx, roomID, userID, role)
}
