package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID       uuid.UUID      `gorm:"type:uuid;primaryKey" json:"user_id"`
	Name         string         `gorm:"type:varchar(100);not null" json:"name"`
	Email        string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-" json:"password"`
	JoinedAt     time.Time      `gorm:"autoCreateTime" json:"joined_at"`
	RoomCredits  int            `gorm:"default:1" json:"room_credits"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	//RoomsCreated []Room       `gorm:"foreignKey:CreatedBy"`
	//RoomMembers []RoomMember `gorm:"foreignKey:UserID"`
	//Expenses    []Expense    `gorm:"foreignKey:UserID"`
	//Credits     []Credits    `gorm:"foreignKey:UserID"`
}

//	func (User) TableName() string {
//		return "users"
//	}
//
// /default:CURRENT_TIMESTAMP
// autoCreateTime
// ;default:gen_random_uuid()
func (User) TableName() string {
	return "users"
}
