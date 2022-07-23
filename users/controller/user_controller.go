package controller

import (
	"chatapp/backend/users/service"
	"chatapp/backend/errors"
	"net/http"
	"strconv"
	userMappings "chatapp/backend/users"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserById(c *gin.Context)
	GetUserByEmailAndUsername(c *gin.Context)
}

type UserControllerImpl struct {
	router *gin.Engine
	userService service.UserService
}

func NewUserController(router *gin.Engine, userService service.UserService) UserController {
	userController := UserControllerImpl{router: router, userService: userService}
	userController.router.GET("/users/:id", userController.GetUserById)
	userController.router.GET("/users", userController.GetUserByEmailAndUsername)
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


// only one querystring parameter allowed. this does not support having both, email and username, as the querystring param. in the service it will be rejected with a custom invalid input error.
func (controller *UserControllerImpl) GetUserByEmailAndUsername(c *gin.Context) {
  
}