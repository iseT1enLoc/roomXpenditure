package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	RoomID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"room_id"`
	RoomName  string    `gorm:"type:varchar(100);not null" json:"room_name"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	Creator     User         `gorm:"foreignKey:CreatedBy"`
	RoomMembers []RoomMember `gorm:"foreignKey:RoomID"`
	Expenses    []Expense    `gorm:"foreignKey:RoomID"`
}

func (Room) TableName() string {
	return "rooms"
}
