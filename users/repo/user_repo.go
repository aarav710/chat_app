package repo

import (
	"chatapp/backend/ent"
	"chatapp/backend/ent/login"
	"chatapp/backend/ent/user"
	"context"
)


type UserRepo interface {
	GetUserById(id int) *ent.User
	GetUserByEmail(email string) *ent.User
	GetUsersContainingUsername(username string) []*ent.User
}

type UserRepoImpl struct {
	ctx context.Context
	db *ent.Client
}

func NewUserRepo(ctx context.Context, db *ent.Client) UserRepo {
	return &UserRepoImpl{ctx: ctx, db: db}
}

func (repo *UserRepoImpl) GetUserById(id int) *ent.User {
  user, _ := repo.db.User.Query().Where(user.ID(id)).Only(repo.ctx)
	return user
}

func (repo *UserRepoImpl) GetUserByEmail(email string) *ent.User {
	user, _ := repo.db.Login.Query().Where(login.Email(email)).QueryUser().Only(repo.ctx)
	return user
}

func (repo *UserRepoImpl) GetUsersContainingUsername(username string) []*ent.User {
	users, _ := repo.db.Login.Query().Where(login.UsernameContains(username)).QueryUser().All(repo.ctx)
	return users
}