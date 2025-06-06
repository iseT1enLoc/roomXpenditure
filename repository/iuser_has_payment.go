package repository

import (
	"703room/703room.com/models"
	"context"

	"github.com/google/uuid"
)

type UserHashPaymentRepository interface {
	CreateExpense(ctx context.Context, userhaspayment *models.UserHasPayment) error

	GetExpensesFiltered(ctx context.Context, userID uuid.UUID, room_id uuid.UUID, year, month, day string) ([]models.UserHasPayment, error)
	GetExpenseByUserID(ctx context.Context, userid uuid.UUID, room_id uuid.UUID) (*models.Expense, error)

	CalculateMemberExpenseByMemberId(ctx context.Context, userID uuid.UUID, room_id uuid.UUID, year, month, day string) (float64, error)
}
