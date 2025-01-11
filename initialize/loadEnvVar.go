package initialize

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVar() {
	// Getting my .env Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
