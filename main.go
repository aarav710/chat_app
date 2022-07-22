package main

import (
	"chatapp/backend/db"
	"chatapp/backend/ent/migrate"
	"chatapp/backend/middleware"
	"context"
	"log"

	userController "chatapp/backend/users/controller"
	userRepo "chatapp/backend/users/repo"
	userService "chatapp/backend/users/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.ErrorHandler)
	db, err := db.Open()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer db.Close()
  ctx := context.Background()
	if err := db.Schema.Create(ctx, migrate.WithDropIndex(true),migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	
// INITIALIZING REPOS
  userRepo := userRepo.NewUserRepo(ctx, db)

// INITIALIZING SERVICES

  userService := userService.NewUserService(userRepo)

// INITIALIZING CONTROLLERS
	userController.NewUserController(router, userService)
	router.Run()
}
