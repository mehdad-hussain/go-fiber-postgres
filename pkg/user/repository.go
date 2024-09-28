package user

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByToken(token string) (*User, error) {
	var user User
	if err := r.DB.Where("activation_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ActivateUser(user *User) error {
	user.Activated = true
	return r.DB.Save(user).Error
}

// GetUserByEmail retrieves a user by their email address.
func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// In user_repository.go
func (repo *UserRepository) GetUserIDByEmail(email string) (uint, error) {
	var user User
	err := repo.DB.Select("id").Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return user.ID, nil
}
