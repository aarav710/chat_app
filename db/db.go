package db

import (
	"chatapp/backend/ent"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Open() (*ent.Client, error) {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }

	dbConnectionString := "postgres://aaravjain:@localhost:5432/chatapp?sslmode=disable"
	client, err := ent.Open("postgres", dbConnectionString)
	return client, err
}