package contacts

import (
	"errors"

	"gorm.io/gorm"
)

type ContactRepository struct {
	DB *gorm.DB
}

// NewContactRepository creates a new contact repository.
func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{DB: db}
}

// GetAllContacts retrieves all contacts for a user with pagination.
func (repo *ContactRepository) GetAllContacts(page, limit int) ([]Contact, error) {
	var contacts []Contact
	err := repo.DB.
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&contacts).Error
	return contacts, err
}

// CreateContact adds a new contact for the user.
func (repo *ContactRepository) CreateContact(contact *Contact) error {
	return repo.DB.Create(contact).Error
}

// UpdateContact updates a contact's details.
func (repo *ContactRepository) UpdateContact(contact *Contact) error {
	return repo.DB.Save(contact).Error
}

// DeleteContact deletes a contact by ID.
func (repo *ContactRepository) DeleteContact(id uint) error {
	return repo.DB.Delete(&Contact{}, id).Error
}

func (repo *ContactRepository) GetContactByEmail(email string) (*Contact, error) {
	var contact Contact
	err := repo.DB.Where("email = ?", email).First(&contact).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &contact, nil
}
