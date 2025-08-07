package api

import (
	"log"

	database "github.com/MatheusBenetti/one-piece-api/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env")
	}

	database.Connect()
}
