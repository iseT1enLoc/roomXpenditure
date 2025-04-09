package models

import (
	"time"

	"github.com/google/uuid"
)

type Credits struct {
	PaymentID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"credit_id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	CreditsAdded  int       `gorm:"not null" json:"credits_added"`
	PaymentMethod string    `gorm:"type:varchar(50)" json:"payment_method"`
	PaidAt        time.Time `gorm:"autoCreateTime" json:"paid_at"`
	Status        string    `gorm:"type:varchar(50);default:'completed'" json:"status"`

	User User `gorm:"foreignKey:UserID"`
}
