package serviceimpl

import (
	"context"
	"log"

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
