package service

import (
	"chatapp/backend/auth"
	"chatapp/backend/ent"
	"chatapp/backend/ent/login"
	"chatapp/backend/login/repo"
)

type LoginService interface {
	CreateUserLogin(password, username, email string) (*ent.Login, error)
}

type LoginServiceImpl struct {
	loginRepo   repo.LoginRepo
	authService auth.AuthService
}

func NewLoginService(loginRepo repo.LoginRepo, authService auth.AuthService) LoginService {
	LoginService := LoginServiceImpl{loginRepo: loginRepo, authService: authService}
	return &LoginService
}

func (service *LoginServiceImpl) CreateUserLogin(password, username, email string) (*ent.Login, error) {
	userRecord, err := service.authService.CreateUser(username, email, password)
	if err != nil {
		return nil, err
	}
	setUserClaimsErr := make(chan error)

	go func() {
		claims := make(map[string]interface{})
        claims["status"] = login.StatusINCOMPLETE_REGISTRATION
        err := service.authService.SetUserClaims(userRecord.UID, claims)
		setUserClaimsErr <- err
	}()

	login, err := service.loginRepo.CreateUser(login.StatusINCOMPLETE_REGISTRATION, userRecord.DisplayName, userRecord.Email, userRecord.UID)
	if err != nil {
		return nil, err
	}

	err = <- setUserClaimsErr
	if err != nil {
		return nil, err
	}
	
	return login, err
}
