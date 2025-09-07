package services

import (
	"context"

	"703room/703room.com/models"

	"github.com/google/uuid"
)

type UserHasPaymentService interface {
	CreateExpense(ctx context.Context, userhaspayment *models.UserHasPayment) error
	GetExpensesFiltered(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, year, month, day string) ([]models.UserHasPayment, error)
	GetExpenseByUserID(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) (*models.Expense, error)
	CalculateMemberExpenseByMemberId(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, year, month, day string) (float64, error)
	GetRoomExpenseDetails(ctx context.Context, roomID uuid.UUID, year, month, day string) ([]models.UserPaymentResponse, error)
}
