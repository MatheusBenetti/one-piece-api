package api

import (
	"log"

	"github.com/MatheusBenetti/one-piece-api/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env")
	}

	database.Connect()
}
