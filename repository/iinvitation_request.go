package repository

import (
	"703room/703room.com/models"
	"context"

	"github.com/google/uuid"
)

type IInvitationManagement interface {
	GetAllPendingInvitationByUserId(ctx context.Context, userID uuid.UUID) ([]models.RoomExpenseInvitationRecipient, error)
	CreateInvitationWithRecipients(ctx context.Context, invitation *models.RoomExpenseInvitationRequest, recipients []models.RoomExpenseInvitationRecipient) error
	UpdateInvitationRequest(ctx context.Context, recipientID uuid.UUID, status models.InvitationStatus) error
	GetRecipientWithInvitation(ctx context.Context, recipientID uuid.UUID) (*models.RoomExpenseInvitationRecipient, error)
}
