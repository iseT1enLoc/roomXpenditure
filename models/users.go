package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID       uuid.UUID      `gorm:"type:uuid;primaryKey;unique" json:"user_id"`
	Name         string         `gorm:"type:varchar(100);not null" json:"name"`
	Email        string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-" json:"password"`
	JoinedAt     time.Time      `gorm:"autoCreateTime" json:"joined_at"`
	RoomCredits  int            `gorm:"default:1" json:"room_credits"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	GoogleId     string         `gorm:"type:varchar(255);unique;not null" json:"google_id"` // New Google ID field

	// Many-to-many via UserHasPayment
	Payments []UserHasPayment `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"payments,omitempty"`
	Rooms    []Room           `gorm:"many2many:room_members"`
}

type ReadableUser struct {
	UserID      uuid.UUID `gorm:"type:uuid;primaryKey;unique" json:"user_id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Email       string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	JoinedAt    time.Time `gorm:"autoCreateTime" json:"joined_at"`
	RoomCredits int       `gorm:"default:1" json:"room_credits"`
	GoogleId    string    `gorm:"type:varchar(255);unique;not null" json:"google_id"` // New Google ID field
}

func (User) TableName() string {
	return "users"
}
