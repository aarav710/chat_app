package controller

import (
	"chatapp/backend/auth"
	"chatapp/backend/errors"
	authenticationController "chatapp/backend/middleware"
	"net/http"
	"strconv"

	chatMappings "chatapp/backend/chats"
	"chatapp/backend/chats/service"
	userService "chatapp/backend/users/service"
	userMappings "chatapp/backend/users"

	"github.com/gin-gonic/gin"
)


type ChatControllerImpl struct {
  router *gin.Engine
  chatService service.ChatService
  authenticationController.AuthenticationController
  userService userService.UserService
}

type ChatController interface {
  FindChatsByUser(c *gin.Context)
  FindChatById(c *gin.Context)
  CreateChat(c *gin.Context)
  UpdateChat(c *gin.Context)
  DeleteChat(c *gin.Context)
  AddUserToChat(c *gin.Context)
  RemoveUserFromChat(c *gin.Context)
}

func NewChatController(router *gin.Engine, chatService service.ChatService, authService auth.AuthService, userService userService.UserService) ChatController {
	authenticationController := authenticationController.NewAuthenticationController(router, authService)
	chatController := ChatControllerImpl{chatService: chatService, router: router, AuthenticationController: authenticationController, userService: userService}
	chatController.router.GET("/users/:userId/chats", chatController.AuthorizeUser([]string{auth.ROLE_USER}), chatController.FindChatsByUser)
	chatController.router.GET("/chats/:chatId", chatController.AuthorizeUser([]string{auth.ROLE_USER}), chatController.FindChatById)
	chatController.router.POST("/chats", chatController.AuthorizeUser([]string{auth.ROLE_USER}), chatController.CreateChat)
	chatController.router.PUT("/chats/:chatId", chatController.AuthorizeUser([]string{auth.ROLE_USER}), chatController.UpdateChat)
	chatController.router.DELETE("/chats/:chatId", chatController.AuthorizeUser([]string{auth.ROLE_USER}), chatController.DeleteChat)
	chatController.router.POST("/chats/:chatId/users/:userId", chatController.AuthorizeUser([]string{auth.ROLE_USER}), chatController.AddUserToChat)
	chatController.router.DELETE("/chats/:chatId/users/:userId", chatController.AuthorizeUser([]string{auth.ROLE_USER}), chatController.RemoveUserFromChat)
	return &chatController
}

func (controller *ChatControllerImpl) FindChatsByUser(c *gin.Context) {
	uid := c.GetString("uid")
    user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	chats, err := controller.chatService.FindChatsByUserId(user.ID)
	if err != nil {
		c.Error(err)
		return 
	}
    var chatsResponse []chatMappings.ChatResponse
	for _, chat := range chats {
		chatResponse := chatMappings.EntToChatResponse(chat)
		chatsResponse = append(chatsResponse, chatResponse)
	}
	c.JSON(http.StatusOK, chatsResponse)
}


func (controller *ChatControllerImpl) FindChatById(c *gin.Context) {
	uid := c.GetString("uid")
	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	chatIdParam := c.Param("chatId")
	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	chat, err := controller.chatService.FindChatById(user, chatId)
	if err != nil {
		c.Error(err)
		return
	}
	chatResponse := chatMappings.EntToChatResponse(chat)
	c.JSON(http.StatusOK, chatResponse)
}

func (controller *ChatControllerImpl) CreateChat(c *gin.Context) {
    uid := c.GetString("uid")
	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	var chatRequest chatMappings.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		c.Error(err)
		return
	}
	chat, err := controller.chatService.CreateChat(user, chatRequest)
    if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, chat)
}

func (controller *ChatControllerImpl) UpdateChat(c *gin.Context) {
    uid := c.GetString("uid")
	chatIdParam := c.Param("chatId")
	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	var chatRequest chatMappings.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		c.Error(err)
		return
	}
	chat, err := controller.chatService.UpdateChat(chatRequest, chatId, uid)
    if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, chat)
}

func (controller *ChatControllerImpl) DeleteChat(c *gin.Context) {
	uid := c.GetString("uid")
    chatIdParam := c.Param("chatId")
	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	err = controller.chatService.DeleteChat(user, chatId)
	if err != nil {
		c.Error(err)
		return 
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (controller *ChatControllerImpl) AddUserToChat(c *gin.Context) {
	uid := c.GetString("uid")
	userIdParam := c.Param("userId")
	toBeAddedUserId, err := strconv.Atoi(userIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	chatIdParam := c.Param("chatId")
	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	addedUser, err := controller.chatService.AddUserToChat(user, toBeAddedUserId, chatId)
	if err != nil {
		c.Error(err)
		return
	}
	addedUserResponse, err := userMappings.EntToResponse(addedUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, addedUserResponse)
}

func (controller *ChatControllerImpl) RemoveUserFromChat(c *gin.Context) {
	uid := c.GetString("uid")
	userIdParam := c.Param("userId")
	toBeRemovedUserId, err := strconv.Atoi(userIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	chatIdParam := c.Param("chatId")
	chatId, err := strconv.Atoi(chatIdParam)
	if err != nil {
		c.Error(errors.InvalidNumericParameterInputError)
		return
	}
	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		c.Error(err)
		return
	}
	err = controller.chatService.RemoveUserFromChat(user, toBeRemovedUserId, chatId)
	if err != nil {
		c.Error(err)
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}