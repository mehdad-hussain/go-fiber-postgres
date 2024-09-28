package main

import (
	"log"

	"github.com/mehdad-hussain/go-fiber-postgres/internal/db" // Import your db package
	"github.com/mehdad-hussain/go-fiber-postgres/pkg/user"    // Import the user model
)

func main() {
	// Connect to the database using the existing Connect function
	db.Connect()

	// Run AutoMigrate for User model using the already established connection
	err := db.DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully.")
}
