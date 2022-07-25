package repo

import (
	"chatapp/backend/ent"
	"chatapp/backend/ent/login"
	"chatapp/backend/ent/user"
	"context"
)


type UserRepo interface {
	GetUserById(id int) (*ent.User, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetUsersContainingUsername(username string) ([]*ent.User, error)
	GetUserByUid(uid string) (*ent.User, error)
}

type UserRepoImpl struct {
	ctx context.Context
	db *ent.Client
}

func NewUserRepo(ctx context.Context, db *ent.Client) UserRepo {
	return &UserRepoImpl{ctx: ctx, db: db}
}

func (repo *UserRepoImpl) GetUserById(id int) (*ent.User, error) {
  user, err := repo.db.User.Query().Where(user.ID(id)).Only(repo.ctx)
	return user, err
}

func (repo *UserRepoImpl) GetUserByEmail(email string) (*ent.User, error) {
	user, err := repo.db.Login.Query().Where(login.Email(email)).QueryUser().Only(repo.ctx)
	return user, err
}

func (repo *UserRepoImpl) GetUsersContainingUsername(username string) ([]*ent.User, error) {
	users, err := repo.db.Login.Query().Where(login.UsernameContains(username)).QueryUser().All(repo.ctx)
	return users, err
}

func (repo *UserRepoImpl) GetUserByUid(uid string) (*ent.User, error) {
	user, err := repo.db.Login.Query().Where(login.UID(uid)).QueryUser().Only(repo.ctx)
	return user, err
}