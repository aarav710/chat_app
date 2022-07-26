package service

import (
	"chatapp/backend/auth"
	"chatapp/backend/ent"
	"chatapp/backend/login/repo"
)

type LoginService interface {
	CreateUserLogin(password, username, email string) (*ent.Login, error)
}

type LoginServiceImpl struct {
	loginRepo repo.LoginRepo
	authService auth.AuthService
}

func NewUserService(loginRepo repo.LoginRepo, authService auth.AuthService) LoginService {
	LoginService := LoginServiceImpl{loginRepo: loginRepo, authService: authService}
	return &LoginService
}

//add role to the database too and setting of claims etc
func (service *LoginServiceImpl) CreateUserLogin(password, username, email string) (*ent.Login, error) {
  userRecord, err := service.authService.CreateUser(username, email, password)
	if err != nil {
		return nil, err
	}
	login, err := service.loginRepo.CreateUser(userRecord.DisplayName, userRecord.Email, userRecord.UID)
	if err != nil {
		return nil, err
	}
	return login, err
}
