package contacts

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/mehdad-hussain/go-fiber-postgres/pkg/user"
)

// Initialize contact handler
var contactRepo *ContactRepository
var userRepo *user.UserRepository

// InitializeContactHandler initializes the contact handler with a repository.
func InitializeContactHandler(contactRepository *ContactRepository, userRepository *user.UserRepository) {
	contactRepo = contactRepository
	userRepo = userRepository
}

// GetContacts retrieves all contacts for the authenticated user.
func GetContacts(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
	}

	contacts, err := contactRepo.GetAllContacts(page, limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to retrieve contacts"})
	}

	return c.JSON(contacts)
}

// CreateContact handles the creation of a contact
func CreateContact(c *fiber.Ctx) error {
	req := new(Contact)
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userID, err := userRepo.GetUserIDByEmail(req.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user ID"})
	}

	if userID == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	req.UserID = userID

	existingContact, err := contactRepo.GetContactByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check for existing contact"})
	}
	if existingContact != nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "A contact with this email already exists"})
	}

	if err := contactRepo.CreateContact(req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create contact"})
	}

	return c.Status(http.StatusCreated).JSON(req)
}

// UpdateContact updates the contact information.
func UpdateContact(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid contact ID"})
	}

	userID, _ := c.Locals("user_id").(uint) // Retrieve the user ID from the middleware

	contact := &Contact{}
	if err := c.BodyParser(contact); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	contact.ID = uint(id)
	contact.UserID = userID // Ensure the user ID is set

	if err := contactRepo.UpdateContact(contact); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update contact"})
	}

	return c.Status(http.StatusOK).JSON(contact)
}

// DeleteContact deletes a specific contact by ID.
func DeleteContact(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid contact ID"})
	}

	if err := contactRepo.DeleteContact(uint(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete contact"})
	}

	return c.Status(http.StatusNoContent).SendString("")
}
