package contacts

import "gorm.io/gorm"

// Contact represents a user's contact.
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	UserID uint   `json:"user_id"`
}
