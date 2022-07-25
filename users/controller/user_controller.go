package controller

import (
	"chatapp/backend/errors"
	userMappings "chatapp/backend/users"
	"chatapp/backend/users/service"
	goError "errors"
	"net/http"
	"strconv"

	"chatapp/backend/auth"
	authenticationController "chatapp/backend/middleware"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserById(c *gin.Context)
	GetUsersUsername(c *gin.Context)
	Me(c *gin.Context)
}

type UserControllerImpl struct {
	router      *gin.Engine
	userService service.UserService
	authenticationController.AuthenticationController
}

func NewUserController(router *gin.Engine, userService service.UserService, authService auth.AuthService) UserController {
	authenticationController := authenticationController.NewAuthenticationController(router, authService)
	userController := UserControllerImpl{router: router, userService: userService, AuthenticationController: authenticationController}
	userController.router.GET("/users/:id", authenticationController.AuthorizeUser([]string{auth.ROLE_USER}), userController.GetUserById)
	userController.router.GET("/users", userController.GetUsersUsername)
	userController.router.GET("/users/me", authenticationController.AuthorizeUser([]string{auth.ROLE_USER}), userController.Me)
	return &userController
}

func (controller *UserControllerImpl) GetUserById(c *gin.Context) {
	userIdParam := c.Param("id")
	id, err := strconv.Atoi(userIdParam)
	if err != nil {
		c.Error(errors.NewInvalidNumericParameterInputError())
		return
	}
	user, err := controller.userService.GetUserById(id)
	if err != nil {
		c.Error(err)
		return
	}
	userResponse := userMappings.EntToResponse(user)
	c.JSON(http.StatusOK, userResponse)
}

func (controller *UserControllerImpl) GetUsersUsername(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.Error(errors.NewIncorrectQueryParameterError("provide a part of the username you would like to search for. an empty string is not valid"))
		return
	}
	users, err := controller.userService.GetUsersByUsername(username)
	if err != nil {
		c.Error(err)
	}
	var usersResponse []userMappings.UserResponse
	for _, user := range users {
		userResponse := userMappings.EntToResponse(user)
		usersResponse = append(usersResponse, userResponse)
	}
	c.JSON(http.StatusOK, usersResponse)
}

func (controller *UserControllerImpl) Me(c *gin.Context) {
	uid, uidExists := c.Get("uid")
	if !uidExists {
		c.Error(goError.New("unexpected failure"))
		return
	}
	switch uid := uid.(type) {
	case string:
		user, err := controller.userService.GetUserByUid(uid)
    if err != nil {
			c.Error(err)
			return
		}
		userResponse := userMappings.EntToResponse(user)
		c.JSON(http.StatusOK, userResponse)
	default:
		c.Error(goError.New("type assertion failure"))
		return
	}
}
