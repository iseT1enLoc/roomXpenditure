package serviceimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"
	"context"
	"errors"
	"fmt"
	"log"
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
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Signup registers a new user.
func (s *authService) Signup(ctx context.Context, user *models.User) (string, error) {
	// Check for invalid user data
	if user == nil || user.Email == "" || user.PasswordHash == "" || user.Name == "" {
		return "", errors.New("invalid user data")
	}

	// Check if the email is already registered
	existing, _ := s.userRepo.GetByEmail(ctx, user.Email)
	if existing != nil {
		return "", errors.New("email already registered")
	}

	// Hash the plain password (not the hashed one)
	hashedPwd, err := s.HashPassword(user.PasswordHash)
	if err != nil {
		return "", err
	}

	// Save the hashed password
	user.PasswordHash = hashedPwd
	user.UserID = uuid.New()

	// Save the user to the database
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		fmt.Println(err)
		return "", err
	}

	// Generate JWT token for the user
	token, err := s.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login authenticates a user and returns a token.
func (s *authService) Login(ctx context.Context, email, password string) (*models.User, string, error) {
	// Retrieve the user from the database by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		// Don't log password mismatch errors for security reasons
		return nil, "", errors.New("invalid email or password")
	}

	// Log that we are comparing password (you can remove this in production)
	fmt.Println("Stored hash:", user.PasswordHash)
	fmt.Println("Input password:", password)

	// Compare the hashed password with the user's input password
	if !s.CheckPasswordHash(password, user.PasswordHash) {
		return nil, "", errors.New("invalid email or password")
	}

	// Generate a JWT token for the user if password matches
	log.Println(user)
	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	// Return the user and the generated token
	return user, token, nil
}

// GenerateToken creates a JWT token for a user.
func (s *authService) GenerateToken(user *models.User) (string, error) {
	log.Println(user.UserID)
	claims := jwt.MapClaims{
		"user_id": user.UserID,
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

// HashPassword hashes plain password
func (s *authService) HashPassword(password string) (string, error) {
	// Hash the password using bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares password with hash
func (s *authService) CheckPasswordHash(password, hash string) bool {
	// Compare the entered password (plain-text) with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// No need to log the error, we just return true/false based on the result
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// ValidateToken verifies a JWT token and returns user info.
func ValidateToken(tokenStr string) (*string, *uuid.UUID, error) {
	fmt.Println(tokenStr)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		log.Println("JWT parse error or token is not valid:", err)
		return nil, nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("JWT claims casting failed")
		return nil, nil, errors.New("invalid claims")
	}

	// Extract and validate email
	emailVal, ok := claims["email"]
	if !ok {
		log.Println("JWT does not contain email")
		return nil, nil, errors.New("email not found in token")
	}
	email, ok := emailVal.(string)
	if !ok || email == "" {
		log.Println("JWT email is not a string or is empty")
		return nil, nil, errors.New("invalid email in token")
	}

	// Extract and parse user_id
	userIDVal, ok := claims["user_id"]
	if !ok {
		log.Println("JWT does not contain user_id")
		return nil, nil, errors.New("user_id not found in token")
	}
	userIDStr, ok := userIDVal.(string)
	if !ok {
		log.Println("JWT user_id is not a string")
		return nil, nil, errors.New("invalid user_id in token")
	}
	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		log.Println("Failed to parse user_id as UUID:", err)
		return nil, nil, errors.New("invalid UUID format")
	}

	fmt.Println("Extracted from token:", email, uid)
	return &email, &uid, nil
}
