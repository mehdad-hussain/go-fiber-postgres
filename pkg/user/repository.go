package user

import (
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
