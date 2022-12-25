package users

import (
	authService "chatapp/backend/auth"
	"chatapp/backend/ent"
	"chatapp/backend/ent/migrate"
	loginRepo "chatapp/backend/login/repo"
	"chatapp/backend/middleware"
	userController "chatapp/backend/users/controller"
	userRepo "chatapp/backend/users/repo"
	userService "chatapp/backend/users/service"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	_ "github.com/lib/pq" 

)

func TestUser(t *testing.T) {
	router := gin.Default()
	creds := option.WithCredentialsFile("../.././service-account-key-file.json")
	router.Use(middleware.ErrorHandler)
	db, err := ent.Open("postgres", "postgres://aaravjain:@localhost:5432/chatapp?sslmode=disable")
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
	userRepo := userRepo.NewUserRepo(ctx, db)
	loginRepo := loginRepo.NewLoginRepo(ctx, db)
	userService := userService.NewUserService(userRepo, loginRepo, authService)
	userController.NewUserController(router, userService, authService)
	
	w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodGet, "/users/hello", nil)
    router.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Errorf("wrong response")
	}
}