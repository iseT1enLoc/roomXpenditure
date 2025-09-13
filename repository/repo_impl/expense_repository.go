package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) repository.ExpenseRepository {
	return &expenseRepository{db: db}
}

// GetExpenseByUserID implements repository.ExpenseRepository.
func (r *expenseRepository) GetExpenseByUserID(ctx context.Context, user_id uuid.UUID, filter_mode []string) ([]models.Expense, error) {
	var expenses_of_one_user []models.Expense
	if err := r.db.WithContext(ctx).Where("user_id = ?", user_id).Find(&expenses_of_one_user).Error; err != nil {
		return nil, err
	}
	return expenses_of_one_user, nil
}

// CreateExpense inserts a new expense record into the database.
func (r *expenseRepository) CreateExpense(ctx context.Context, expense *models.Expense) error {
	log.Println("CREATED EXPENSE")
	return r.db.WithContext(ctx).Create(expense).Error
}

// GetExpenseByID fetches an expense by its ID.
func (r *expenseRepository) GetExpenseByID(ctx context.Context, id uuid.UUID) (*models.Expense, error) {
	var expense models.Expense
	if err := r.db.WithContext(ctx).First(&expense, "expense_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

// ListExpensesByRoomID retrieves all expenses for a given room.
func (r *expenseRepository) ListExpensesByRoomID(ctx context.Context, roomID string) ([]models.Expense, error) {
	var expenses []models.Expense
	if err := r.db.WithContext(ctx).Where("room_id = ?", roomID).Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

// DeleteExpense removes an expense by its ID.
func (r *expenseRepository) DeleteExpense(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("expense_id = ?", id).Delete(&models.Expense{}).Error
}

// UpdateExpense updates an existing expense record.
func (r *expenseRepository) UpdateExpense(ctx context.Context, expense *models.Expense) error {
	return r.db.WithContext(ctx).Save(expense).Error
}
func (r *expenseRepository) GetExpensesFiltered(
	ctx context.Context,
	userID uuid.UUID,
	year, month, day string,
) ([]models.Expense, error) {
	var expenses []models.Expense
	query := r.db.WithContext(ctx).Where("user_id = ?", userID)
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

// CalculateMemberExpenseByMemberId implements repository.ExpenseRepository.
func (r *expenseRepository) CalculateMemberExpenseByMemberId(ctx context.Context, userID uuid.UUID, year, month, day string) (float64, error) {
	var expenses []models.Expense
	query := r.db.WithContext(ctx).Where("user_id = ?", userID)
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

// GetExpensesFilteredFromStartDateToEndDate implements repository.UserHashPaymentRepository.
func (u *expenseRepository) GetExpensesFilteredFromStartDateToEndDate(ctx context.Context, userID, roomID uuid.UUID, start_date *time.Time, end_date *time.Time) ([]models.UserHasPayment, error) {
	var expenses []models.UserHasPayment
	query := u.db.WithContext(ctx).Where("user_id = ? AND room_id = ?", userID, roomID)
	if start_date != nil && end_date != nil {
		query = query.Where("used_date BETWEEN ? AND ?", *start_date, *end_date)
	} else if start_date != nil {
		query = query.Where("used_date >= ?", start_date)
	} else if end_date != nil {
		query = query.Where("used_date<= ?", end_date)
	}
	err := query.Order("used_date DESC").Find(&expenses).Error
	return expenses, err
}
