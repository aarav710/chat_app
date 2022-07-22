package service

import (
	"chatapp/backend/ent"
	"chatapp/backend/users/repo"
)

type UserService interface {
	GetUserById(userId int) *ent.User
}

type UserServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	userService := UserServiceImpl{userRepo: userRepo}
	return &userService
}

func (service *UserServiceImpl) GetUserById(userId int) *ent.User {
  user := service.userRepo.GetUserById(userId)
	return user
}

