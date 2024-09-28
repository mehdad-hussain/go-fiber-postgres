package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	JWTSecret string
	DBUrl     string
)

func LoadConfig() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}

	// Construct the path to the .env file
	envPath := filepath.Join(dir, "config", ".env")

	// Load the .env file
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Load the environment variables
	JWTSecret = os.Getenv("JWT_SECRET")
	DBUrl = os.Getenv("DATABASE_URL")
}
