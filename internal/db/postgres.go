package db

import (
	"log"

	"github.com/mehdad-hussain/go-fiber-postgres/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	config.LoadConfig()

	var err error
	DB, err = gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Database connection established.")

}
