package serviceimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"703room/703room.com/services"
	"context"
	"errors"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(
	userRepo repository.UserRepository,
) services.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// RegisterUser registers a new user
func (s *userService) RegisterUser(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	return s.userRepo.CreateUser(ctx, user)
}

// GetUserByID retrieves a user by their UUID string
func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return s.userRepo.GetByID(ctx, id)
}

// GetUserByEmail retrieves a user by email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email is empty")
	}
	return s.userRepo.GetByEmail(ctx, email)
}

// UpdateUser updates user details
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return s.userRepo.Update(ctx, user)
}

// DeleteUser soft-deletes a user by ID
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is empty")
	}
	return s.userRepo.Delete(ctx, id)
}

// GetAllUsers implements services.UserService.
func (s *userService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}
