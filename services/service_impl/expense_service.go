package serviceimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"
	"context"
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
)

type expenseService struct {
	expenseRepo repository.ExpenseRepository
}

func NewExpenseService(expenseRepo repository.ExpenseRepository) services.ExpenseService {
	return &expenseService{
		expenseRepo: expenseRepo,
	}
}

// GetListExpensesByUserID implements services.ExpenseService.
func (s *expenseService) GetExpensesFiltered(
	ctx context.Context,
	userID uuid.UUID,
	year, month, day string,
) ([]models.Expense, error) {
	return s.expenseRepo.GetExpensesFiltered(ctx, userID, year, month, day)
}

// CreateExpense adds a new expense to a room.
func (s *expenseService) CreateExpense(ctx context.Context, expense *models.Expense) error {
	log.Println("EXPENSE SERVICE CREATED")
	if expense == nil {
		return errors.New("expense cannot be nil")
	}

	if expense.Amount <= 0 {
		return errors.New("expense amount must be greater than 0")
	}
	if strings.TrimSpace(expense.Notes) == "" {
		return errors.New("expense description is required")
	}
	if expense.UserID == uuid.Nil {
		return errors.New("created_by must be a valid UUID")
	}
	log.Println("EXPENSE SERVICE CREATED")
	return s.expenseRepo.CreateExpense(ctx, expense)
}

// GetExpenseByID returns a single expense by its ID.
func (s *expenseService) GetExpenseByID(ctx context.Context, id uuid.UUID) (*models.Expense, error) {
	if id == uuid.Nil {
		return nil, errors.New("expense ID is required")
	}
	return s.expenseRepo.GetExpenseByID(ctx, id)
}

// ListExpensesByRoomID returns all expenses in a specific room.
func (s *expenseService) ListExpensesByRoomID(ctx context.Context, roomID string) ([]models.Expense, error) {
	if roomID == "" {
		return nil, errors.New("room ID is required")
	}
	return s.expenseRepo.ListExpensesByRoomID(ctx, roomID)
}

// DeleteExpense deletes a specific expense by ID.
func (s *expenseService) DeleteExpense(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("expense ID is required")
	}
	return s.expenseRepo.DeleteExpense(ctx, id)
}

// UpdateExpense updates an existing expense.
func (s *expenseService) UpdateExpense(ctx context.Context, expense *models.Expense) error {
	if expense == nil {
		return errors.New("expense cannot be nil")
	}
	if expense.ExpenseID == uuid.Nil {
		return errors.New("expense ID must be a valid UUID")
	}
	if expense.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	if strings.TrimSpace(expense.Notes) == "" {
		return errors.New("description cannot be empty")
	}
	return s.expenseRepo.UpdateExpense(ctx, expense)
}

// CalculateMemberExpenseByMemberId implements services.ExpenseService.
func (s *expenseService) CalculateMemberExpenseByMemberId(ctx context.Context, userID uuid.UUID, year string, month string, day string) (float64, error) {
	total_money, err := s.expenseRepo.CalculateMemberExpenseByMemberId(ctx, userID, year, month, day)
	if err != nil {
		return 0, err
	}
	return total_money, nil

}
