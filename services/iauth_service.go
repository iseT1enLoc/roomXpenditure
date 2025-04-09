package services

import (
	"703room/703room.com/models"
	"context"
)

type AuthService interface {
	GenerateToken(user *models.User) (string, error)
	ValidateToken(tokenStr string) (*models.User, error)
	Login(ctx context.Context, email, password string) (*models.User, string, error)
	Signup(ctx context.Context, user *models.User) (string, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
