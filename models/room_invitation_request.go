package models

import (
	"time"

	"github.com/google/uuid"
)

type InvitationStatus string

const (
	InvitationPending  InvitationStatus = "pending"
	InvitationAccepted InvitationStatus = "accepted"
	InvitationDenied   InvitationStatus = "denied"
)

type RoomExpenseInvitationRequest struct {
	ID                uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FromUserId        *uuid.UUID `json:"from_user_id" gorm:"type:uuid"`
	RoomId            uuid.UUID  `json:"room_id" gorm:"type:uuid;not null;index"`
	InvitationMessage string     `json:"invitation_message"`
	CreatedAt         time.Time  `json:"created_at" gorm:"autoCreateTime"`

	// Nếu user bị xóa → FromUserId = NULL (không mất invitation)
	FromUser *User `gorm:"foreignKey:FromUserId;constraint:OnDelete:SET NULL"`
	Room     *Room `gorm:"foreignKey:RoomId;constraint:OnDelete:SET NULL"`

	Recipients []RoomExpenseInvitationRecipient `json:"recipients,omitempty" gorm:"foreignKey:InvitationId;constraint:OnDelete:CASCADE"`
}
