package middleware

import (
	"chatapp/backend/auth"
	"net/http"

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

/*
func (controller *AuthenticationControllerImpl) AuthorizeUser(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("authorization")
		if authToken == "" || !strings.HasPrefix(authToken, "Bearer ") {
            c.Error(errors.UnauthorizedError)
			return
		}
        fmt.Println("yo lmao")
		splitToken := strings.Split(authToken, "Bearer ")
		if len(splitToken) != 2 {
			c.Error(errors.UnauthorizedError)
			return
		}
		authToken = splitToken[1]
		uid, claims, err := controller.authService.VerifyUserIdToken(authToken)
		if err != nil {
			c.Error(err)
			return
		}
		for _, role := range roles {
			if claims[role] == nil {
				c.Error(errors.UnauthorizedError)
				return
			}
		}
		c.Set("uid", uid)
		c.Next()
	}
}
*/

func (controller *AuthenticationControllerImpl) AuthorizeUser(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("uid", "ckuEB1qhPbT1Ks96yg7w2qY7p333")
		c.Next()
	}
}

func AuthorizeWebsocketConnection(request *http.Request) string {
	// actual implementation is left
	return "ckuEB1qhPbT1Ks96yg7w2qY7p333"
}