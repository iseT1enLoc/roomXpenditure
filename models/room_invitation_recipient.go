package models

import "github.com/google/uuid"

type RoomExpenseInvitationRecipient struct {
	ID           uuid.UUID        `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	InvitationId uuid.UUID        `json:"invitation_id" gorm:"type:uuid;not null;index"`
	UserId       *uuid.UUID       `json:"user_id" gorm:"type:uuid"`
	Status       InvitationStatus `json:"status" gorm:"type:varchar(20);default:'pending'"`
	// Nếu user bị xóa → record recipient cũng bị xoá theo
	User       *User                        `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Invitation RoomExpenseInvitationRequest `json:"invitation,omitempty" gorm:"foreignKey:InvitationId"`
}
