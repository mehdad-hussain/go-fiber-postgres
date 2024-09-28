package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string `json:"name"`
	Email           string `json:"email" gorm:"unique"`
	PasswordHash    string `json:"-"` // Exclude this from JSON response
	Activated       bool   `json:"activated"`
	ActivationToken string `json:"activation_token"`
}
