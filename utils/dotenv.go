package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnv(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error loading .env file")
	}

	return os.Getenv(key)
}