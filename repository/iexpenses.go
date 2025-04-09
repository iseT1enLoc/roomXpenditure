package repository

import (
	"703room/703room.com/models"
	"context"
)

type ExpenseRepository interface {
	CreateExpense(ctx context.Context, expense *models.Expense) error
	GetExpenseByID(ctx context.Context, id string) (*models.Expense, error)
	ListExpensesByRoomID(ctx context.Context, roomID string) ([]models.Expense, error)
	DeleteExpense(ctx context.Context, id string) error
	UpdateExpense(ctx context.Context, expense *models.Expense) error
}
