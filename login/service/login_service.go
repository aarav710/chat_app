package service

import (
	"chatapp/backend/auth"
	"chatapp/backend/ent/login"
	"chatapp/backend/login/repo"
)

type LoginService interface {
	CreateUserLogin(password, username, email string) (string, error)
}

type LoginServiceImpl struct {
	loginRepo   repo.LoginRepo
	authService auth.AuthService
}

func NewLoginService(loginRepo repo.LoginRepo, authService auth.AuthService) *LoginServiceImpl {
	LoginService := LoginServiceImpl{loginRepo: loginRepo, authService: authService}
	return &LoginService
}

func (service *LoginServiceImpl) CreateUserLogin(password, username, email string) (string, error) {
	userRecord, err := service.authService.CreateUser(username, email, password)
	if err != nil {
		return "", err
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
		return "", err
	}

	err = <- setUserClaimsErr
	if err != nil {
		return "", err
	}
	jwt, err := service.authService.CreateAuthToken(login.UID)
	if err != nil {
		return "", err
	}
	return jwt, err
}