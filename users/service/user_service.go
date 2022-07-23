package service

import (
	"chatapp/backend/ent"
	"chatapp/backend/users/repo"
)

type UserService interface {
	GetUserById(userId int) (*ent.User, error)
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

