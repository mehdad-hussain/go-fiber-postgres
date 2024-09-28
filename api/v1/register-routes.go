package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdad-hussain/go-fiber-postgres/internal/handlers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1") // Create a group with the /api/v1 prefix

	api.Post("/users", handlers.CreateUser)            // POST /api/v1/users
	api.Post("/users/activate", handlers.ActivateUser) // POST /api/v1/users/activate
	// api.Post("/token/auth", handlers.AuthenticateUser)  // POST /api/v1/token/auth

	// Add the new GET route
	api.Get("/health", handlers.HealthCheck) // GET /api/v1/health
}
