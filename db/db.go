package db

import (
	"chatapp/backend/ent"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Open() (*ent.Client, error) {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	client, err := ent.Open("postgres", dbConnectionString)
	return client, err
}