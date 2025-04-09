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

type expenseService struct {
	expenseRepo repository.ExpenseRepository
}

func NewExpenseService(expenseRepo repository.ExpenseRepository) services.ExpenseService {
	return &expenseService{
		expenseRepo: expenseRepo,
	}
}

// CreateExpense adds a new expense to a room.
func (s *expenseService) CreateExpense(ctx context.Context, expense *models.Expense) error {
	if expense == nil {
		return errors.New("expense cannot be nil")
	}
	// if expense.RoomID == uuid.Nil {
	// 	return errors.New("room ID must be a valid UUID")
	// }
	if expense.Amount <= 0 {
		return errors.New("expense amount must be greater than 0")
	}
	if strings.TrimSpace(expense.Notes) == "" {
		return errors.New("expense description is required")
	}
	if expense.UserID == uuid.Nil {
		return errors.New("created_by must be a valid UUID")
	}
	return s.expenseRepo.CreateExpense(ctx, expense)
}

// GetExpenseByID returns a single expense by its ID.
func (s *expenseService) GetExpenseByID(ctx context.Context, id string) (*models.Expense, error) {
	if id == "" {
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
