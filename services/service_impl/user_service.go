package serviceimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"
	"context"
	"errors"

	"github.com/google/uuid"
)

type userService struct {
	userRepo       repository.UserRepository
	roomRepo       repository.RoomRepository
	roomMemberRepo repository.RoomMemberRepository
	expenseRepo    repository.ExpenseRepository
}

// AddExpense allows a user to add an expense to a room
func (s *userService) AddExpense(ctx context.Context, expense *models.Expense) error {
	if expense == nil {
		return errors.New("expense is nil")
	}
	if expense.RoomID == uuid.Nil || expense.UserID == uuid.Nil || expense.Title == "" {
		return errors.New("invalid expense data")
	}
	return s.expenseRepo.CreateExpense(ctx, expense)
}

// UpdateExpense allows updating an existing expense
func (s *userService) UpdateExpense(ctx context.Context, expense *models.Expense) error {
	if expense == nil || expense.ExpenseID == uuid.Nil {
		return errors.New("invalid expense")
	}
	return s.expenseRepo.UpdateExpense(ctx, expense)
}

// DeleteExpense deletes an expense by its ID
func (s *userService) DeleteExpense(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("expense id is empty")
	}
	return s.expenseRepo.DeleteExpense(ctx, id)
}

// GetExpensesByRoomID returns all expenses for a room
func (s *userService) GetExpensesByRoomID(ctx context.Context, roomID string) ([]models.Expense, error) {
	if roomID == "" {
		return nil, errors.New("room id is empty")
	}
	return s.expenseRepo.ListExpensesByRoomID(ctx, roomID)
}

func NewUserService(
	userRepo repository.UserRepository,
	roomRepo repository.RoomRepository,
	roomMemberRepo repository.RoomMemberRepository,
) services.UserService {
	return &userService{
		userRepo:       userRepo,
		roomRepo:       roomRepo,
		roomMemberRepo: roomMemberRepo,
	}
}
func (s *userService) CreateRoom(ctx context.Context, userID string, roomName string) (*models.Room, error) {
	room := &models.Room{
		RoomName:  roomName,
		CreatedBy: uuid.MustParse(userID),
	}

	if err := s.roomRepo.Create(ctx, room); err != nil {
		return nil, err
	}

	// Auto join as admin
	member := &models.RoomMember{
		RoomID: room.RoomID,
		UserID: uuid.MustParse(userID),
		Role:   "admin",
	}
	if err := s.roomMemberRepo.AddMember(ctx, member); err != nil {
		return nil, err
	}

	return room, nil
}

func (s *userService) JoinRoom(ctx context.Context, userID, roomID string) error {
	member := &models.RoomMember{
		RoomID: uuid.MustParse(roomID),
		UserID: uuid.MustParse(userID),
		Role:   "member",
	}
	return s.roomMemberRepo.AddMember(ctx, member)
}

func (s *userService) LeaveRoom(ctx context.Context, userID, roomID string) error {
	return s.roomMemberRepo.RemoveMember(ctx, roomID, userID)
}

func (s *userService) ListUserRooms(ctx context.Context, userID string) ([]models.Room, error) {
	return s.roomRepo.ListByUserID(ctx, userID)
}

// RegisterUser registers a new user
func (s *userService) RegisterUser(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return s.userRepo.Create(ctx, user)
}

// GetUserByID retrieves a user by their UUID string
func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return s.userRepo.GetByID(ctx, id)
}

// GetUserByEmail retrieves a user by email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email is empty")
	}
	return s.userRepo.GetByEmail(ctx, email)
}

// UpdateUser updates user details
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return s.userRepo.Update(ctx, user)
}

// DeleteUser soft-deletes a user by ID
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is empty")
	}
	return s.userRepo.Delete(ctx, id)
}
