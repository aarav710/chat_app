package auth

import (
	"context"

	auth "firebase.google.com/go/v4/auth"
)


type AuthService interface {
	//returns uid
	VerifyUserIdToken(authToken string) (*string, map[string]interface{}, error)
	//SetUserClaims()
}

type AuthServiceImpl struct {
	ctx context.Context
	AuthClient *auth.Client
}

func NewAuthService(authClient *auth.Client, ctx context.Context) AuthService {
  return &AuthServiceImpl{AuthClient: authClient, ctx: ctx}
}

func (AuthService *AuthServiceImpl) VerifyUserIdToken(authToken string) (*string, map[string]interface{}, error) {
  token, err := AuthService.AuthClient.VerifyIDToken(AuthService.ctx, authToken)
	if err != nil {
    return nil, nil, err
	}
	return &token.UID, token.Claims, err
}
