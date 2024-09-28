package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mehdad-hussain/go-fiber-postgres/pkg/user"
	"golang.org/x/crypto/bcrypt"
)

var userRepo *user.UserRepository

// InitializeUserHandler initializes the user handler with a repository.
func InitializeUserHandler(repository *user.UserRepository) {
	if repository == nil {
		log.Fatal("UserRepository is not initialized")
	}
	userRepo = repository
}

// Helper function to generate an activation token.
func generateToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "API is working!"})
}

// CreateUser creates a new user and returns an activation token.
func CreateUser(c *fiber.Ctx) error {
	if userRepo == nil {
		log.Fatal("UserRepository is not initialized")
	}

	req := new(struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error hashing password"})
	}

	// Create user
	newUser := &user.User{
		Name:            req.Name,
		Email:           req.Email,
		PasswordHash:    string(hashedPassword),
		ActivationToken: generateToken(),
	}

	// Create the user in the database
	if err := userRepo.CreateUser(newUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message":         "User created successfully",
		"activationToken": newUser.ActivationToken,
	})
}

// ActivateUser activates a user's account with the given token.
func ActivateUser(c *fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Token is required"})
	}

	user, err := userRepo.GetUserByToken(token)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Invalid token"})
	}

	if err := userRepo.ActivateUser(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to activate user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User activated successfully"})
}
