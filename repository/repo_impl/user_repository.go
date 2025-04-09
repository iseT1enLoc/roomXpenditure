package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Create implements repository.UserRepository.
func (u *userRepository) Create(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Create(&user).Error
}

// Delete implements repository.UserRepository.
func (u *userRepository) Delete(ctx context.Context, id string) error {
	return u.db.Delete(&models.User{}, "user_id = ?", id).Error
}

// GetByEmail implements repository.UserRepository.
func (u *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID implements repository.UserRepository.
func (u *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	if err := u.db.Where("user_id", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update implements repository.UserRepository.
func (u *userRepository) Update(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}
