package services

import (
	"703room/703room.com/models"
)

type AuthService interface {
	GenerateToken(user *models.User) (string, error)
	ValidateToken(tokenStr string) (*models.User, error)
}
