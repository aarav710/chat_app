package controller

import (
	"chatapp/backend/auth"
	"chatapp/backend/errors"
	messageMappings "chatapp/backend/messages"
	"chatapp/backend/messages/service"
	authenticationController "chatapp/backend/middleware"
	"net/http"
	"strconv"
	"chatapp/backend/messages/hub"
	userService "chatapp/backend/users/service"
	"github.com/gin-gonic/gin"
)

type MessageController interface {
	GetMessagesByChatId(c *gin.Context)
	CreateMessage(c *gin.Context)
}

type MessageControllerImpl struct {
	router      *gin.Engine
	messageService service.MessageService
	userService userService.UserService
	authenticationController.AuthenticationController
	hub  hub.Hub
}

func NewMessageController(router *gin.Engine, messageService service.MessageService, authService auth.AuthService, hub hub.Hub, userService userService.UserService) MessageController {
	authenticationController := authenticationController.NewAuthenticationController(router, authService)
	messageController := MessageControllerImpl{router: router, messageService: messageService, AuthenticationController: authenticationController, hub: hub, userService: userService}
	messageController.router.GET("/chats/:chatId/messages", messageController.AuthorizeUser([]string{auth.ROLE_USER}), messageController.GetMessagesByChatId)
	messageController.router.POST("/chats/:chatId/messages", messageController.AuthorizeUser([]string{auth.ROLE_USER}), messageController.CreateMessage)
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
	var entityToResponseErr error
	for _, message := range messages {
		messageResponse, err := messageMappings.EntToResponse(message)
		if err != nil {
            entityToResponseErr = err
			break
		}
		messagesResponse = append(messagesResponse, messageResponse)
	}
	if entityToResponseErr != nil {
		c.Error(err)
		return
	}
    c.JSON(http.StatusOK, messagesResponse)
}

func (controller *MessageControllerImpl) CreateMessage(c *gin.Context) {
	uid := c.GetString("uid")
	chatIdParam := c.Param("chatId")
	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	var messageRequest messageMappings.MessageRequest
	if err := c.ShouldBindJSON(&messageRequest); err != nil {
		c.Error(err)
		return
	}
	message, err := controller.messageService.CreateMessage(messageRequest, chatId, uid)
	if err != nil {
		c.Error(err)
	}
	messageResponse, err := messageMappings.EntToResponse(message)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusCreated, messageResponse)
}