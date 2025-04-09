package serviceimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) services.AuthService {
	return &authService{userRepo: userRepo}
}

// JWT secret (ideally from env)
var jwtSecret = []byte(os.Getenv("JWT_SECRJWT_SECRETET"))

// Signup registers a new user.
func (s *authService) Signup(ctx context.Context, user *models.User) (string, error) {
	if user == nil || user.Email == "" || user.PasswordHash == "" || user.Name == "" {
		return "", errors.New("invalid user data")
	}

	existing, _ := s.userRepo.GetByEmail(ctx, user.Email)
	if existing != nil {
		return "", errors.New("email already registered")
	}

	hashedPwd, err := s.HashPassword(user.PasswordHash)
	if err != nil {
		return "", err
	}

	user.PasswordHash = hashedPwd
	user.UserID = uuid.New()
	fmt.Println(user)
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		fmt.Println(err)
		return "", err
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login authenticates a user and returns token.
func (s *authService) Login(ctx context.Context, email, password string) (*models.User, string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, "", errors.New("invalid email or password")
	}

	if !s.CheckPasswordHash(password, user.PasswordHash) {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// GenerateToken creates a JWT token for a user.
func (s *authService) GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.UserID.String(),
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken verifies a JWT token and returns user info.
func (s *authService) ValidateToken(tokenStr string) (*models.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid email in token")
	}

	ctx := context.Background()
	return s.userRepo.GetByEmail(ctx, email)
}

// HashPassword hashes the given plain password.
func (s *authService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares plain password with hash.
func (s *authService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
