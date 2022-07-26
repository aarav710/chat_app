package main

import (
	"chatapp/backend/db"
	"chatapp/backend/ent/migrate"
	"chatapp/backend/middleware"
	"context"
	"log"

	authService "chatapp/backend/auth"

	userController "chatapp/backend/users/controller"
	userRepo "chatapp/backend/users/repo"
	userService "chatapp/backend/users/service"

	loginController "chatapp/backend/login/controller"
	loginRepo "chatapp/backend/login/repo"
  loginService "chatapp/backend/login/service"

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

	authService := authService.NewAuthService(auth, ctx)
	

	// INITIALIZING REPOS
	userRepo := userRepo.NewUserRepo(ctx, db)
	loginRepo := loginRepo.NewUserRepo(ctx, db)

	// INITIALIZING SERVICES

	userService := userService.NewUserService(userRepo)
	loginService := loginService.NewUserService(loginRepo, authService)

	// INITIALIZING CONTROLLERS
	userController.NewUserController(router, userService, authService)
	loginController.NewLoginController(router, loginService)
	router.Run()
}
