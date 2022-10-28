package controller

import (
	"chatapp/backend/auth"
	"chatapp/backend/errors"
	messageMappings "chatapp/backend/messages"
	"chatapp/backend/messages/service"
	authenticationController "chatapp/backend/middleware"
	"net/http"
	"strconv"

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
	messageController.router.GET("/chats/:chatId/messages", messageController.AuthorizeUser([]string{auth.ROLE_USER}), messageController.GetMessagesByChatId)
	return &messageController
}

func (controller *MessageControllerImpl) GetMessagesByChatId(c *gin.Context) {
	uid := c.GetString("uid")
	chatIdParam := c.Param("chatId")
	cursorQuery := c.DefaultQuery("cursor", "0")
    cursor, err := strconv.Atoi(cursorQuery)
	if err != nil {
		c.Error(err)
		return
	}
	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	messages, err := controller.messageService.FindMessagesByChatId(uid, chatId, cursor)
	if err != nil {
		c.Error(err)
	}
	var messagesResponse []messageMappings.MessageResponse
	var entityToResponse error
	for _, message := range messages {
		messageResponse, err := messageMappings.EntToResponse(message)
		if err != nil {
            entityToResponse = err
			break
		}
		messagesResponse = append(messagesResponse, messageResponse)
	}
	if entityToResponse != nil {
		c.Error(err)
		return
	}
    c.JSON(http.StatusOK, messagesResponse)
}