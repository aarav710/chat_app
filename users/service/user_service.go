package service

import (
	"chatapp/backend/ent"
	"chatapp/backend/users/repo"
)

type UserService interface {
	GetUserById(userId int) (*ent.User, error)
	GetUsersByUsername(username string) ([]*ent.User, error)
	GetUserByUid(uid string) (*ent.User, error)
}

type UserServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	userService := UserServiceImpl{userRepo: userRepo}
	return &userService
}

func (service *UserServiceImpl) GetUserById(userId int) (*ent.User, error) {
  user, err := service.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (service *UserServiceImpl) GetUsersByUsername(username string) ([]*ent.User, error) {
	users, err := service.userRepo.GetUsersContainingUsername(username)
	if err != nil {
		return nil, err
	}
	return users, err
}

func (service *UserServiceImpl) GetUserByUid(uid string)(*ent.User, error) {
	user, err := service.userRepo.GetUserByUid(uid)
	if err != nil {
		return nil, err
	}
	return user, err
}