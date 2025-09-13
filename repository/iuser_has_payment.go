package repository

import (
	"703room/703room.com/models"
	"context"
	"time"

	"github.com/google/uuid"
)

type UserHashPaymentRepository interface {
	CreateExpense(ctx context.Context, userhaspayment *models.UserHasPayment) error

	GetExpensesFiltered(ctx context.Context, userID uuid.UUID, room_id uuid.UUID, year, month, day string) ([]models.UserHasPayment, error)
	GetExpenseByUserID(ctx context.Context, userid uuid.UUID, room_id uuid.UUID) (*models.Expense, error)

	CalculateMemberExpenseByMemberId(ctx context.Context, userID uuid.UUID, room_id uuid.UUID, year, month, day string) (float64, error)
	GetRoomExpenseDetails(ctx context.Context, room_id uuid.UUID, year, month, day string) ([]models.UserPaymentResponse, error)

	GetExpensesFilteredFromStartDateToEndDate(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, start_date, end_date *time.Time) ([]models.UserPaymentResponse, error)
}
