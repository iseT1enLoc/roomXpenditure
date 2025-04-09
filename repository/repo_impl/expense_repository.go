package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"

	"gorm.io/gorm"
)

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) repository.ExpenseRepository {
	return &expenseRepository{db: db}
}

// CreateExpense inserts a new expense record into the database.
func (r *expenseRepository) CreateExpense(ctx context.Context, expense *models.Expense) error {
	return r.db.WithContext(ctx).Create(expense).Error
}

// GetExpenseByID fetches an expense by its ID.
func (r *expenseRepository) GetExpenseByID(ctx context.Context, id string) (*models.Expense, error) {
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
