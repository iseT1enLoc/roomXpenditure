package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type userhaspaymentRepository struct {
	db *gorm.DB
}

// CalculateMemberExpenseByMemberId implements repository.UserHashPaymentRepository.
func (u *userhaspaymentRepository) CalculateMemberExpenseByMemberId(ctx context.Context, userID uuid.UUID, room_id uuid.UUID, year string, month string, day string) (float64, error) {
	var expenses []models.UserHasPayment
	query := u.db.WithContext(ctx).Where("user_id = ? AND room_id=?", userID, room_id)
	if year != "" {
		query = query.Where("EXTRACT(YEAR FROM created_at) = ?", year)
	}
	if month != "" {
		query = query.Where("EXTRACT(MONTH FROM created_at) = ?", month)
	}
	if day != "" {
		query = query.Where("EXTRACT(DAY FROM created_at) = ?", day)
	}

	err := query.Order("created_at DESC").Find(&expenses).Error

	var total_expense float64
	for i := 0; i < len(expenses); i = i + 1 {
		total_expense = total_expense + expenses[i].Amount
	}
	return total_expense, err
}

func (u *userhaspaymentRepository) CreateExpense(ctx context.Context, userhaspayment *models.UserHasPayment) error {
	// Ensure the ID is set
	if userhaspayment.ID == uuid.Nil {
		userhaspayment.ID = uuid.New()
	}

	// Set the creation time
	userhaspayment.CreatedAt = time.Now()

	// Create the record in the database
	return u.db.WithContext(ctx).Create(userhaspayment).Error
}

// GetExpenseByUserID implements repository.UserHashPaymentRepository.
func (u *userhaspaymentRepository) GetExpenseByUserID(ctx context.Context, userid uuid.UUID, room_id uuid.UUID) (*models.Expense, error) {
	var expense models.Expense
	if err := u.db.WithContext(ctx).First(&expense, "user_id = ? AND room_id ", userid, room_id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

// GetExpensesFiltered implements repository.UserHashPaymentRepository.
func (u *userhaspaymentRepository) GetExpensesFiltered(ctx context.Context, userID uuid.UUID, room_id uuid.UUID, year string, month string, day string) ([]models.UserHasPayment, error) {
	var expenses []models.UserHasPayment
	query := u.db.WithContext(ctx).Where("user_id = ? AND room_id = ?", userID, room_id)
	if year != "" {
		query = query.Where("EXTRACT(YEAR FROM created_at) = ?", year)
	}
	if month != "" {
		query = query.Where("EXTRACT(MONTH FROM created_at) = ?", month)
	}
	if day != "" {
		query = query.Where("EXTRACT(DAY FROM created_at) = ?", day)
	}

	err := query.Order("created_at DESC").Find(&expenses).Error
	return expenses, err
}

func NewUserHasPaymentRepository(db *gorm.DB) repository.UserHashPaymentRepository {
	return &userhaspaymentRepository{
		db: db,
	}
}
