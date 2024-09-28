package initializer

import (
	"log"

	"github.com/mehdad-hussain/go-fiber-postgres/internal/db"
	"github.com/mehdad-hussain/go-fiber-postgres/internal/handlers"
	"github.com/mehdad-hussain/go-fiber-postgres/pkg/user"
)

// Initialize initializes the repositories and handlers.
func Initialize() {
	// Initialize UserRepository and pass it to handlers
	userRepository := user.NewUserRepository(db.DB)
	handlers.InitializeUserHandler(userRepository)

	// You can initialize other repositories and handlers here as needed.
	log.Println("Repositories and handlers initialized successfully.")
}
