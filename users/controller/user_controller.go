package controller

import (
	"chatapp/backend/errors"
	userMappings "chatapp/backend/users"
	"chatapp/backend/users/service"
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
	UpdateUser(c *gin.Context)
	FindUsersByChatId(c *gin.Context)
	CreateUser(c *gin.Context)
}

type UserControllerImpl struct {
	router      *gin.Engine
	userService service.UserService
	authenticationController.AuthenticationController
}

func NewUserController(router *gin.Engine, userService service.UserService, authService auth.AuthService) *UserControllerImpl {
	authenticationController := authenticationController.NewAuthenticationController(router, authService)
	userController := UserControllerImpl{router: router, userService: userService, AuthenticationController: authenticationController}
	userController.router.GET("/users/:userId", userController.AuthorizeUser([]string{auth.ROLE_USER}), userController.GetUserById)
	userController.router.GET("/users", userController.GetUsersUsername)
	userController.router.GET("/users/me", userController.AuthorizeUser([]string{auth.ROLE_USER}), userController.Me)
	userController.router.PUT("/users", userController.AuthorizeUser([]string{auth.ROLE_USER}), userController.UpdateUser)
	userController.router.GET("chats/:chatId/users", userController.FindUsersByChatId)
	userController.router.POST("/completeRegistration", userController.AuthorizeUser([]string{auth.ROLE_USER}), userController.CreateUser)
	return &userController
}

func (controller *UserControllerImpl) GetUserById(c *gin.Context) {
	userIdParam := c.Param("userId")
	id, err := strconv.Atoi(userIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	user, err := controller.userService.GetUserById(id)
	if err != nil {
		c.Error(err)
		return
	}
	userResponse, err := userMappings.EntToResponse(user)
	if err != nil {
		c.Error(err)
		return
	}
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
	usersResponse := make([]userMappings.UserResponse, 0)
	for _, user := range users {
		userResponse, err := userMappings.EntToResponse(user)
		if err != nil {
			c.Error(err)
			return
		}
		usersResponse = append(usersResponse, userResponse)
	}
	c.JSON(http.StatusOK, usersResponse)
}

func (controller *UserControllerImpl) Me(c *gin.Context) {
	uid := c.GetString("uid")

	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	userResponse, err := userMappings.EntToResponse(user)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, userResponse)

}

func (controller *UserControllerImpl) UpdateUser(c *gin.Context) {
	uid := c.GetString("uid")
	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	var userRequest userMappings.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.Error(err)
		return
	}
	updatedUser, err := controller.userService.UpdateUser(userRequest, user.ID)
	if err != nil {
		c.Error(err)
		return
	}
	userResponse, err := userMappings.EntToResponse(updatedUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, userResponse)
}

func (controller *UserControllerImpl) FindUsersByChatId(c *gin.Context) {
	chatIdParam := c.Param("chatId")
	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	users, err := controller.userService.FindUsersByChatId(chatId)
	if err != nil {
		c.Error(err)
		return
	}
	usersResponse := make([]userMappings.UserResponse, 0)
	for _, user := range users {
		userResponse, err := userMappings.EntToResponse(user)
		if err != nil {
			c.Error(err)
			return
		}
		usersResponse = append(usersResponse, userResponse)
	}
	c.JSON(http.StatusOK, usersResponse)
}

func (controller *UserControllerImpl) CreateUser(c *gin.Context) {
	uid := c.GetString("uid")
	var userRequest userMappings.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.Error(err)
		return
	}
	user, err := controller.userService.CreateUser(uid, userRequest)
	if err != nil {
		c.Error(err)
		return
	}
	userResponse, err := userMappings.EntToResponse(user)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, userResponse)
}
