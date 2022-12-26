package main

import (
	"chatapp/backend/db"
	"chatapp/backend/ent"
	"chatapp/backend/ent/migrate"

	//"chatapp/backend/messages/hub"
	messagesHub "chatapp/backend/messages/hub"
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


	controllers := InitializeDI(ctx, router, db, auth)
	router.GET("/chats/ws", func (c *gin.Context) {
		// add user authentication for this endpoint over here
		messagesHub.ServeWebSocketRequests(controllers.Hub, c.Writer, c.Request, &ent.User{})
	})
	router.Run()
}

