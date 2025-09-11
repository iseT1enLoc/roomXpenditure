package serviceimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"
	"context"
	"errors"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type roomService struct {
	roomRepo       repository.RoomRepository
	userRepo       repository.UserRepository
	roomMemberRepo repository.RoomMemberRepository
	invitationRepo repository.IInvitationManagement
}

// UpdateMemberCount implements services.RoomService.
func (s *roomService) UpdateMemberCount(ctx context.Context) error {
	rooms, err := s.roomRepo.GetAllRooms(ctx)
	log.Println(len(rooms))
	if err != nil {
		return err
	}
	for i := 0; i < len(rooms); i = i + 1 {
		log.Println(rooms[i].MemBerCount)
		rooms[i].MemBerCount = len(rooms[i].Members)
		log.Println(rooms[i].MemBerCount)
		s.roomRepo.Save(ctx, &rooms[i])
	}
	return nil
}

// GetAllPendingInvitationByUserId implements services.RoomService.
func (s *roomService) GetAllPendingInvitationByUserId(ctx context.Context, userID uuid.UUID) ([]models.RoomExpenseInvitationRecipient, error) {
	return s.invitationRepo.GetAllPendingInvitationByUserId(ctx, userID)
}

func (s *roomService) SendInvitationToUsers(ctx context.Context, fromUserID, roomID uuid.UUID, emails []string, message string) error {
	// 1. Fetch users by email via repository
	users, err := s.userRepo.FindByEmails(ctx, emails)
	if err != nil {
		return err
	}
	log.Println(fromUserID)
	for i := 0; i < len(users); i = i + 1 {
		log.Println(users[i].Email)
		if users[i].UserID.String() == fromUserID.String() {
			return errors.New(users[i].Email + " Can not invite himself!")
		} else if slices.Contains(emails, users[i].Email) == false {
			return errors.New(users[i].Email + " is not in current system, ask him to login to the system, redo all process!")
		}
	}
	if len(users) == 0 {
		return errors.New("no users found for provided emails")
	}

	// 2. Build invitation
	invitation := &models.RoomExpenseInvitationRequest{
		ID:                uuid.New(),
		FromUserId:        &fromUserID,
		RoomId:            roomID,
		InvitationMessage: message,
		CreatedAt:         time.Now(),
	}

	// 3. Build recipients
	var recipients []models.RoomExpenseInvitationRecipient
	for _, u := range users {
		recipients = append(recipients, models.RoomExpenseInvitationRecipient{
			ID:           uuid.New(),
			InvitationId: invitation.ID,
			UserId:       &u.UserID, // ✅ use ID
			Status:       models.InvitationPending,
		})
	}

	// 4. Save using repository (with transaction)
	return s.invitationRepo.CreateInvitationWithRecipients(ctx, invitation, recipients)
}

func (s *roomService) UpdateInvitationRequestStatus(ctx context.Context, recipientID uuid.UUID, status models.InvitationStatus) error {
	// 1. Update status
	if err := s.invitationRepo.UpdateInvitationRequest(ctx, recipientID, status); err != nil {
		return err
	}

	// 2. If accepted → add member
	if status == models.InvitationAccepted {
		recipient, err := s.invitationRepo.GetRecipientWithInvitation(ctx, recipientID)
		if err != nil {
			return err
		}
		room, err := s.roomRepo.GetByID(ctx, recipient.Invitation.RoomId.String())

		roomMember := &models.RoomMember{
			ID:     uuid.New(),
			RoomID: recipient.Invitation.RoomId,
			UserID: *recipient.UserId,
			Role:   string(models.RoomMemberMem),
		}
		//update and save room member count
		room.MemBerCount += 1
		s.roomRepo.Save(ctx, room)
		return s.roomMemberRepo.AddMember(ctx, roomMember)
	}

	return nil
}

func NewRoomService(roomRepo repository.RoomRepository, roomMemberRepo repository.RoomMemberRepository, userRepo repository.UserRepository, invitationRepo repository.IInvitationManagement) services.RoomService {
	return &roomService{
		roomRepo:       roomRepo,
		userRepo:       userRepo,
		roomMemberRepo: roomMemberRepo,
		invitationRepo: invitationRepo,
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
