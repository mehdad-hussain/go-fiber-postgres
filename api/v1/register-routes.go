package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdad-hussain/go-fiber-postgres/internal/middleware"
	"github.com/mehdad-hussain/go-fiber-postgres/pkg/contacts"
	"github.com/mehdad-hussain/go-fiber-postgres/pkg/user"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/health", user.HealthCheck)

	api.Post("/users", user.CreateUser)
	api.Post("/users/activate", user.ActivateUser)
	api.Post("/token/auth", user.AuthenticateUser)

	api.Get("/contacts", middleware.JWTMiddleware, contacts.GetContacts)
	api.Post("/contacts", middleware.JWTMiddleware, contacts.CreateContact)
	api.Patch("/contacts/:id", middleware.JWTMiddleware, contacts.UpdateContact)
	api.Delete("/contacts/:id", middleware.JWTMiddleware, contacts.DeleteContact)
}
