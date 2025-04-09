package models

import (
	"time"

	"github.com/google/uuid"
)

type RoomMember struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	RoomID   uuid.UUID `gorm:"type:uuid;not null;index:idx_room_user,unique" json:"room_id"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;index:idx_room_user,unique" json:"user_id"`
	Role     string    `gorm:"type:varchar(20);not null;default:'member'" json:"role"`
	JoinedAt time.Time `gorm:"autoCreateTime" json:"joined_at"`
}
