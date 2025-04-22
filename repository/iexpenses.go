package repository

import (
	"703room/703room.com/models"
	"context"

	"github.com/google/uuid"
)

type ExpenseRepository interface {
	CreateExpense(ctx context.Context, expense *models.Expense) error

	GetExpensesFiltered(ctx context.Context, userID uuid.UUID, year, month, day string) ([]models.Expense, error)
	GetExpenseByID(ctx context.Context, id uuid.UUID) (*models.Expense, error)

	CalculateMemberExpenseByMemberId(ctx context.Context, userID uuid.UUID, year, month, day string) (float64, error)

	ListExpensesByRoomID(ctx context.Context, roomID string) ([]models.Expense, error)
	DeleteExpense(ctx context.Context, id string) error
	UpdateExpense(ctx context.Context, expense *models.Expense) error
}
