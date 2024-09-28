package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/mehdad-hussain/go-fiber-postgres/api/v1"
	"github.com/mehdad-hussain/go-fiber-postgres/config"
	"github.com/mehdad-hussain/go-fiber-postgres/internal/db"
	"github.com/mehdad-hussain/go-fiber-postgres/internal/initializer"
)

func main() {
	// Load configuration
	config.LoadConfig()
	log.Println("Configuration loaded...")

	// Connect to the database
	db.Connect()

	// Initialize repositories and handlers
	initializer.Initialize() // Call the initializer

	// Initialize Fiber
	app := fiber.New()
	log.Println("Fiber app initialized...")

	// Set up routes
	v1.RegisterRoutes(app)
	log.Println("Routes have been set up...")

	port := ":3000"

	// Start the server
	log.Println("Listening on port", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Println("Server started successfully.")
}
