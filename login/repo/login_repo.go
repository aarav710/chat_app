package repo

import (
	"chatapp/backend/ent"
	"context"
)

type LoginRepo interface {
	CreateUser(username, email, uid string) (*ent.Login, error) 
}

type LoginRepoImpl struct {
	ctx context.Context
	db *ent.Client
}

func NewUserRepo(ctx context.Context, db *ent.Client) LoginRepo {
	return &LoginRepoImpl{ctx: ctx, db: db}
}

func (repo *LoginRepoImpl) CreateUser(username, email, uid string) (*ent.Login, error) {
  login, err := repo.db.Login.Create().SetEmail(email).SetUsername(username).SetUID(uid).Save(repo.ctx)
	return login, err
}