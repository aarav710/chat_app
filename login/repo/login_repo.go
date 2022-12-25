package repo

import (
	"chatapp/backend/ent"
	"chatapp/backend/ent/login"
	"context"
)

type LoginRepo interface {
	CreateUser(status login.Status, username, email, uid string) (*ent.Login, error)
	FindLoginByUid(uid string) (*ent.Login, error)
	UpdateClaim(status login.Status, login *ent.Login) (*ent.Login, error)
}

type LoginRepoImpl struct {
	ctx context.Context
	db  *ent.Client
}

func NewLoginRepo(ctx context.Context, db *ent.Client) *LoginRepoImpl {
	return &LoginRepoImpl{ctx: ctx, db: db}
}

func (repo *LoginRepoImpl) CreateUser(status login.Status, username, email, uid string) (*ent.Login, error) {
	login, err := repo.db.Login.Create().SetEmail(email).SetUsername(username).SetUID(uid).SetStatus(status).Save(repo.ctx)
	return login, err
}

func (repo *LoginRepoImpl) FindLoginByUid(uid string) (*ent.Login, error) {
	login, err := repo.db.Login.Query().Where(login.UID(uid)).Only(repo.ctx)
	return login, err
}

func (repo *LoginRepoImpl) UpdateClaim(status login.Status, login *ent.Login) (*ent.Login, error) {
	updatedLogin, err := login.Update().SetStatus(status).Save(repo.ctx)
	return updatedLogin, err
}