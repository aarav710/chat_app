package auth

import (
	"context"

	auth "firebase.google.com/go/v4/auth"
)


type AuthService interface {
	//returns uid
	VerifyUserIdToken(authToken string) (*string, map[string]interface{}, error)
	//SetUserClaims()
	//UpdateUserClaim()
	CreateUser(username, email, password string) (*auth.UserRecord, error)
}

type AuthServiceImpl struct {
	ctx context.Context
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
