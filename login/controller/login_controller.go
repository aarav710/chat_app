package controller

import (
	"net/http"

	loginMappings "chatapp/backend/login"
	"chatapp/backend/login/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Register(c *gin.Context)
}

type LoginControllerImpl struct {
	router      *gin.Engine
  loginService service.LoginService
}

func NewLoginController(router *gin.Engine, loginService service.LoginService) LoginController {
	loginController := LoginControllerImpl{router: router, loginService: loginService}
  loginController.router.GET("users/register", loginController.Register)
	return &loginController
}

func (controller *LoginControllerImpl) Register(c *gin.Context) {
  var register loginMappings.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.Error(err)
		return
	}
	login, err := controller.loginService.CreateUserLogin(register.Password, register.Username, register.Email)
	if err != nil {
		c.Error(err)
		return
	}
  c.JSON(http.StatusOK, gin.H{"username": login.Username, "email": login.Email})
}