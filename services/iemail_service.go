package services

import (
	"context"

	"github.com/google/uuid"
)

type EmailService interface {
	SendReportToRoomate(ctx context.Context, room_id uuid.UUID, year, month, day, message string) error
}
