package models

import (
	"time"

	"github.com/google/uuid"
)

type UserHasPayment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	RoomID    uuid.UUID `gorm:"type:uuid;not null;index:idx_room_user,unique" json:"room_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_room_user,unique" json:"user_id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Amount    float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Notes     string    `gorm:"type:text" json:"notes"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
