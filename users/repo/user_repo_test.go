package repo

import (
	"context"
	"testing"

	"chatapp/backend/ent"
	"chatapp/backend/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetUserById(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	ctx := context.Background()
	userRepo := NewUserRepo(ctx, client)
	t.Run("no user to be found", func(t *testing.T) {
		user, err := userRepo.GetUserById(5)
		if user != nil || !ent.IsNotFound(err) {
			t.Errorf("Expected no user to be found.")
		}
	})
}

func TestGetUserByEmail(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	ctx := context.Background()
	userRepo := NewUserRepo(ctx, client)
	t.Run("no user to be found", func(t *testing.T) {
		user, err := userRepo.GetUserByEmail("johnx946@gmail.com")
		if user != nil || !ent.IsNotFound(err) {
			t.Errorf("Expected no user to be found.")
		}
	})
}
