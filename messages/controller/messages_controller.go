package controller

import (
	"chatapp/backend/auth"
	"chatapp/backend/messages/service"
	authenticationController "chatapp/backend/middleware"

	"github.com/gin-gonic/gin"
)

type MessageController interface {
	GetMessagesByChatId(c *gin.Context)
}

type MessageControllerImpl struct {
	router      *gin.Engine
	messageService service.MessageService
	authenticationController.AuthenticationController
}

func NewMessageController(router *gin.Engine, messageService service.MessageService, authService auth.AuthService) MessageController {
	authenticationController := authenticationController.NewAuthenticationController(router, authService)
	messageController := MessageControllerImpl{router: router, messageService: messageService, AuthenticationController: authenticationController}
	messageController.router.GET("/chats/:chatId/", messageController.AuthorizeUser([]string{auth.ROLE_USER}), messageController.GetMessagesByChatId)
	return &messageController
}

func (controller *MessageControllerImpl) GetMessagesByChatId(c *gin.Context) {
	
}