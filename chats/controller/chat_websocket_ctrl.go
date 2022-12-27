package controller

import (
	"chatapp/backend/auth"
	"chatapp/backend/messages/hub"
	authenticationController "chatapp/backend/middleware"
	userService "chatapp/backend/users/service"

	"github.com/gin-gonic/gin"
)

type ChatWebsocketCtrlImpl struct {
	hub hub.Hub
	router *gin.Engine
	userService userService.UserService
	authenticationController.AuthenticationController
}

type ChatWebsocketCtrl interface {
	ServeWebsocketRequests(c *gin.Context)
}

func NewChatWebsocketCtrl(hub hub.Hub, router *gin.Engine, userService userService.UserService, authService auth.AuthService) *ChatWebsocketCtrlImpl {
	authenticationController := authenticationController.NewAuthenticationController(router, authService)
	chatWebsocketCtrl := ChatWebsocketCtrlImpl{hub: hub, router: router, userService: userService, AuthenticationController: authenticationController}
	chatWebsocketCtrl.router.GET("/chats/ws", authenticationController.AuthorizeUser([]string{auth.ROLE_USER}), chatWebsocketCtrl.ServeWebsocketRequests)
	return &chatWebsocketCtrl
}

func (controller *ChatWebsocketCtrlImpl) ServeWebsocketRequests(c *gin.Context) {
	uid := c.GetString("uid")
	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	hub.ServeWebSocketRequests(controller.hub, c.Writer, c.Request, user)
}