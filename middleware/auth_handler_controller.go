package middleware

import (
	"chatapp/backend/auth"
	"chatapp/backend/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenticationController interface {
	AuthorizeUser(roles []string) gin.HandlerFunc
}

type AuthenticationControllerImpl struct {
	router      *gin.Engine
	authService auth.AuthService
}

func NewAuthenticationController(router *gin.Engine, authService auth.AuthService) AuthenticationController {
	return &AuthenticationControllerImpl{router: router, authService: authService}
}

func (controller *AuthenticationControllerImpl) AuthorizeUser(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("authorization")
		splitToken := strings.Split(authToken, "Bearer ")
		authToken = splitToken[1]
		uid, claims, err := controller.authService.VerifyUserIdToken(authToken)
		if err != nil {
			c.Error(err)
			return
		}
		for _, role := range roles {
			if claims[role] == nil {
				c.Error(errors.NewUnauthorizedError())
				return
			}
		}
		c.Set("uid", uid)
		c.Next()
	}
}
