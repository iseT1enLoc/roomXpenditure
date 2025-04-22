package repoimpl

import (
	"703room/703room.com/models"
	"703room/703room.com/repository"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// GetAllUsers implements repository.UserRepository.
func (u *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser implements repository.UserRepository.
func (u *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	tx := u.db.Begin()
	err := tx.Create(&user).Error
	fmt.Println(user)
	if err != nil {
		tx.Rollback()
		return err
	}
	u.db.Debug()

	tx.Commit()

	fmt.Printf("User created successfully with ID: %s\n", user.UserID)
	return nil
}

// Delete implements repository.UserRepository.
func (u *userRepository) Delete(ctx context.Context, id string) error {
	return u.db.Delete(&models.User{}, "user_id = ?", id).Error
}

// GetByEmail implements repository.UserRepository.
func (u *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(user.Name)
	return &user, nil

}

// GetByID implements repository.UserRepository.
func (u *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	err := u.db.WithContext(ctx).First(&user, "user_id = ?", id).Error
	if err != nil {
		// You can check for specific error types if needed
		fmt.Printf("Error fetching user by ID %s: %v\n", id, err)
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
