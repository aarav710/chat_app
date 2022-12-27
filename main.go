package main

import (
	"chatapp/backend/db"
	"chatapp/backend/ent/migrate"

	"chatapp/backend/middleware"

	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	router := gin.Default()
	creds := option.WithCredentialsFile("./service-account-key-file.json")
	router.Use(middleware.ErrorHandler)
	db, err := db.Open()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer db.Close()
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, creds)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	if err := db.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}


	InitializeDI(ctx, router, db, auth)
	router.Run()
}

