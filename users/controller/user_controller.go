package controller

import (
	"chatapp/backend/errors"
	"chatapp/backend/users/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	getUserById(c *gin.Context)
	getUserByEmailAndUsername(c *gin.Context)
}

type UserControllerImpl struct {
	router *gin.Engine
	userService service.UserService
}

func NewUserController(router *gin.Engine, userService service.UserService) UserController {
	userController := UserControllerImpl{router: router, userService: userService}
	userController.router.GET("/users/:id", userController.getUserById)
	userController.router.GET("/users", userController.getUserByEmailAndUsername)
	return &userController
}

func (controller *UserControllerImpl) getUserById(c *gin.Context) {
  userIdParam := c.Param("id")
	id, err := strconv.Atoi(userIdParam)
	if err != nil {
		c.Error(errors.NewNotFoundError("Invalid input parameter type. Provide a integer instead."))
		return
	}
	user := controller.userService.GetUserById(id)
	userResponse := user
	c.JSON(http.StatusOK, userResponse)
}


// only one querystring parameter allowed. this does not support having both, email and username, as the querystring param. in the service it will be rejected with a custom invalid input error.
func (controller *UserControllerImpl) getUserByEmailAndUsername(c *gin.Context) {
  
}