package main

import (
	"log"

	"github.com/mehdad-hussain/go-fiber-postgres/internal/db"
	"github.com/mehdad-hussain/go-fiber-postgres/pkg/contacts"
	"github.com/mehdad-hussain/go-fiber-postgres/pkg/user"
)

func main() {
	// Connect to the database using the existing Connect function
	db.Connect()

	// Run AutoMigrate for both User and Contact models using the established connection
	err := db.DB.AutoMigrate(&user.User{}, &contacts.Contact{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully.")
}
