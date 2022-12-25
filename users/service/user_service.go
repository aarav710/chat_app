package service

import (
	"chatapp/backend/auth"
	"chatapp/backend/ent"
	loginRepo "chatapp/backend/login/repo"
	userMappings "chatapp/backend/users"
	"chatapp/backend/users/repo"
	loginStatus "chatapp/backend/ent/login"

	
)

type UserService interface {
	GetUserById(userId int) (*ent.User, error)
	GetUsersByUsername(username string) ([]*ent.User, error)
	GetUserByUid(uid string) (*ent.User, error)
	UpdateUser(userRequest userMappings.UserRequest, userId int) (*ent.User, error)
	FindUsersByChatId(chatId int)([]*ent.User, error)
	CreateUser(uid string, userRequest userMappings.UserRequest) (*ent.User, error)
}

type UserServiceImpl struct {
	userRepo repo.UserRepo
	loginRepo loginRepo.LoginRepo
	authService auth.AuthService
}

func NewUserService(userRepo repo.UserRepo, loginRepo loginRepo.LoginRepo, authService auth.AuthService) *UserServiceImpl {
	userService := UserServiceImpl{userRepo: userRepo, loginRepo: loginRepo, authService: authService}
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

func (service *UserServiceImpl) GetUserByUid(uid string) (*ent.User, error) {
	user, err := service.userRepo.GetUserByUid(uid)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (service *UserServiceImpl) UpdateUser(userRequest userMappings.UserRequest, userId int) (*ent.User, error) {
	user, err := service.userRepo.UpdateUser(userRequest, userId)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (service *UserServiceImpl) FindUsersByChatId(chatId int)([]*ent.User, error) {
	user, err := service.userRepo.FindUsersByChatId(chatId)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (service *UserServiceImpl) CreateUser(uid string, userRequest userMappings.UserRequest) (*ent.User, error) {
	login, err := service.loginRepo.FindLoginByUid(uid)
	if err != nil {
		return nil, err
	}
    user, err := service.userRepo.CreateUser(userRequest, login)
	if err != nil {
		return nil, err
	}
	login, err = service.loginRepo.UpdateClaim(loginStatus.StatusUSER, login)
	if err != nil {
		return nil, err
	}
	service.authService.UpdateClaimToValidUser(uid)
	user.Edges.Login = login
	return user, err
}
