package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type invitationRepo struct {
	db *gorm.DB
}

func NewInvitationRepo(db *gorm.DB) repository.IInvitationManagement {
	return &invitationRepo{db: db}
}

func (r *invitationRepo) GetAllPendingInvitationByUserId(ctx context.Context, userID uuid.UUID) ([]models.RoomExpenseInvitationRecipient, error) {
	var recipients []models.RoomExpenseInvitationRecipient

	err := r.db.WithContext(ctx).
		Where("user_id = ? AND status = ?", userID, models.InvitationPending).
		Preload("Invitation").
		Preload("Invitation.FromUser").
		Preload("Invitation.Room").
		Find(&recipients).Error

	return recipients, err
}
func (r *invitationRepo) CreateInvitationWithRecipients(ctx context.Context, invitation *models.RoomExpenseInvitationRequest, recipients []models.RoomExpenseInvitationRecipient) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(invitation).Error; err != nil {
			return err
		}
		if len(recipients) > 0 {
			if err := tx.Create(&recipients).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *invitationRepo) UpdateInvitationRequest(ctx context.Context, recipientID uuid.UUID, status models.InvitationStatus) error {
	return r.db.WithContext(ctx).
		Model(&models.RoomExpenseInvitationRecipient{}).
		Where("id = ?", recipientID).
		Update("status", status).Error
}

func (r *invitationRepo) GetRecipientWithInvitation(ctx context.Context, recipientID uuid.UUID) (*models.RoomExpenseInvitationRecipient, error) {
	var recipient models.RoomExpenseInvitationRecipient
	if err := r.db.WithContext(ctx).
		Preload("Invitation").
		Where("id = ?", recipientID).
		First(&recipient).Error; err != nil {
		return nil, err
	}
	return &recipient, nil
}
