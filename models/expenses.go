package models

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ExpenseID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"expense_id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Amount    float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Notes     string    `gorm:"type:text" json:"notes"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	//RoomID uuid.UUID `gorm:"type:uuid;not null" json:"room_id"`
}
