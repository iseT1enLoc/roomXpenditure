package serviceimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"
	"context"
)

type creditsService struct {
	creditsRepo repository.CreditsRepository
}

func NewCreditsService(creditsRepo repository.CreditsRepository) services.CreditsService {
	return &creditsService{creditsRepo: creditsRepo}
}

func (s *creditsService) CreatePayment(ctx context.Context, credit *models.Credits) error {
	return s.creditsRepo.CreatePayment(ctx, credit)
}

func (s *creditsService) GetPaymentByID(ctx context.Context, id string) (*models.Credits, error) {
	return s.creditsRepo.GetPaymentByID(ctx, id)
}

func (s *creditsService) ListPaymentsByUserID(ctx context.Context, userID string) ([]models.Credits, error) {
	return s.creditsRepo.ListPaymentsByUserID(ctx, userID)
}

func (s *creditsService) UpdatePaymentStatus(ctx context.Context, id string, status string) error {
	return s.creditsRepo.UpdatePaymentStatus(ctx, id, status)
}
