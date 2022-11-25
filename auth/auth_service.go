package auth

import (
	"chatapp/backend/ent/login"
	"context"

	auth "firebase.google.com/go/v4/auth"
)

type AuthService interface {
	//returns uid
	VerifyUserIdToken(authToken string) (*string, map[string]interface{}, error)
	
	SetUserClaims(uid string, claims map[string]interface{}) error
	UpdateClaimToValidUser(uid string) error
	CreateUser(username, email, password string) (*auth.UserRecord, error)
	CreateAuthToken(uid string) (string, error)
}

type AuthServiceImpl struct {
	ctx        context.Context
	AuthClient *auth.Client
}

func NewAuthService(authClient *auth.Client, ctx context.Context) AuthService {
	return &AuthServiceImpl{AuthClient: authClient, ctx: ctx}
}

func (authService *AuthServiceImpl) VerifyUserIdToken(authToken string) (*string, map[string]interface{}, error) {
	token, err := authService.AuthClient.VerifyIDToken(authService.ctx, authToken)
	if err != nil {
		return nil, nil, err
	}
	return &token.UID, token.Claims, err
}

func (authService *AuthServiceImpl) CreateUser(username, email, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		DisplayName(username).
		Disabled(false)
	user, err := authService.AuthClient.CreateUser(authService.ctx, params)
	return user, err
}

func (authService *AuthServiceImpl) SetUserClaims(uid string, claims map[string]interface{}) error {
	err := authService.AuthClient.SetCustomUserClaims(authService.ctx, uid, claims)
    return err
}

// this function is used to 
func (authService *AuthServiceImpl) UpdateClaimToValidUser(uid string) error {
	user, err := authService.AuthClient.GetUser(authService.ctx, uid)
	if err != nil {
		return err
	}
	claim := user.CustomClaims
	claim["status"] = login.StatusUSER.String()
	err = authService.AuthClient.SetCustomUserClaims(authService.ctx, uid, claim)
	return err
}

func (authService *AuthServiceImpl) CreateAuthToken(uid string) (string, error) {
	token, err := authService.AuthClient.CustomToken(authService.ctx, uid)
	if err != nil {
		return "", err
	}
	return token, nil
}