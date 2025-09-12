package serviceimpl

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"703room/703room.com/models"
	"703room/703room.com/repository"

	"github.com/google/uuid"
)

type userHasPaymentService struct {
	repo repository.UserHashPaymentRepository
}

// Constructor
func NewUserHasPaymentService(repo repository.UserHashPaymentRepository) *userHasPaymentService {
	return &userHasPaymentService{repo: repo}
}

// CreateExpense creates a new expense entry for a user in a room.
func (s *userHasPaymentService) CreateExpense(ctx context.Context, userhaspayment *models.UserHasPayment) error {
	return s.repo.CreateExpense(ctx, userhaspayment)
}

// GetExpensesFiltered fetches expenses for a user in a room, filtered by date if provided.
func (s *userHasPaymentService) GetExpensesFiltered(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, year, month, day string) ([]models.UserHasPayment, error) {
	return s.repo.GetExpensesFiltered(ctx, userID, roomID, year, month, day)
}

// GetExpenseByUserID returns an expense for a user in a room (assuming a single latest or summary one).
func (s *userHasPaymentService) GetExpenseByUserID(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) (*models.Expense, error) {
	return s.repo.GetExpenseByUserID(ctx, userID, roomID)
}

// CalculateMemberExpenseByMemberId calculates total expenses for a user in a room with optional date filtering.
func (s *userHasPaymentService) CalculateMemberExpenseByMemberId(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, year, month, day string) (float64, error) {
	log.Println("Enter line 39")
	return s.repo.CalculateMemberExpenseByMemberId(ctx, userID, roomID, year, month, day)
}
func (s *userHasPaymentService) GetRoomExpenseDetails(ctx context.Context, roomID uuid.UUID, year, month, day string) ([]models.UserPaymentResponse, error) {
	return s.repo.GetRoomExpenseDetails(ctx, roomID, year, month, day)
}
func (s *userHasPaymentService) GetExpenseFromStartDateToEndDate(ctx context.Context, userId, roomId, StartDate, EndDate string) ([]models.UserHasPayment, error) {
	user_id, err := uuid.Parse(userId)
	if err != nil {
		return nil, errors.New("Invalid user id")
	}
	room_id, err := uuid.Parse(roomId)
	if err != nil {
		return nil, errors.New("Invalid room id")
	}
	var start_date *time.Time
	var end_date *time.Time

	if StartDate != "" {
		t, err := time.Parse("02/01/2006", StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start_date format, use YYYY-MM-DD")
		}
		start_date = &t
	}

	if EndDate != "" {
		t, err := time.Parse("02/01/2006", EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end_date format, use YYYY-MM-DD")
		}
		// ensure end_date includes the whole day
		t = t.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
		end_date = &t
	}

	expenses, err := s.repo.GetExpensesFilteredFromStartDateToEndDate(ctx, user_id, room_id, start_date, end_date)
	return expenses, err
}
