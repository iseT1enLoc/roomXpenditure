package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"

	"gorm.io/gorm"
)

type creditsRepository struct {
	db *gorm.DB
}

func NewCreditsRepository(db *gorm.DB) repository.CreditsRepository {
	return &creditsRepository{db: db}
}

// CreatePayment adds a new credit payment record to the database.
func (r *creditsRepository) CreatePayment(ctx context.Context, credit *models.Credits) error {
	return r.db.WithContext(ctx).Create(credit).Error
}

// GetPaymentByID retrieves a credit payment by its ID.
func (r *creditsRepository) GetPaymentByID(ctx context.Context, id string) (*models.Credits, error) {
	var credit models.Credits
	if err := r.db.WithContext(ctx).First(&credit, "credit_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &credit, nil
}

// ListPaymentsByUserID returns all payments made by a specific user.
func (r *creditsRepository) ListPaymentsByUserID(ctx context.Context, userID string) ([]models.Credits, error) {
	var credits []models.Credits
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&credits).Error; err != nil {
		return nil, err
	}
	return credits, nil
}

// UpdatePaymentStatus updates the status of a payment.
func (r *creditsRepository) UpdatePaymentStatus(ctx context.Context, id string, status string) error {
	return r.db.WithContext(ctx).
		Model(&models.Credits{}).
		Where("credit_id = ?", id).
		Update("status", status).Error
}
