package repository

import (
	"703room/703room.com/models"
	"context"
)

type CreditsRepository interface {
	CreatePayment(ctx context.Context, credit *models.Credits) error
	GetPaymentByID(ctx context.Context, id string) (*models.Credits, error)
	ListPaymentsByUserID(ctx context.Context, userID string) ([]models.Credits, error)
	UpdatePaymentStatus(ctx context.Context, id string, status string) error
}
