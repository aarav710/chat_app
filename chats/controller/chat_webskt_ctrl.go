package controller

import (
	"chatapp/backend/auth"
	chatService "chatapp/backend/chats/service"
	"chatapp/backend/ent"
	"chatapp/backend/messages"
	messageService "chatapp/backend/messages/service"
	authenticationController "chatapp/backend/middleware"
	userService "chatapp/backend/users/service"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

type ChatWebsktCtrlImpl struct {
	router *gin.Engine
	userService userService.UserService
	chatService chatService.ChatService
	messageService messageService.MessageService
	m *melody.Melody
}

type ChatWebsktCtrl interface {
	ServeWebsocketRequests(c *gin.Context)
	HandleConnect(s *melody.Session)
	HandleBroadcastingMessage(msg []byte, sessions []*melody.Session) error 
}

func NewChatWebsktCtrl(router *gin.Engine, userService userService.UserService, authService auth.AuthService, chatService chatService.ChatService, messageService messageService.MessageService) *ChatWebsktCtrlImpl {
	m := melody.New()
	chatWebsocketCtrl := ChatWebsktCtrlImpl{router: router, userService: userService, m: m, chatService: chatService, messageService: messageService}
	chatWebsocketCtrl.router.GET("/chats/ws", chatWebsocketCtrl.ServeWebsocketRequests)
	chatWebsocketCtrl.m.HandleConnect(chatWebsocketCtrl.HandleConnect)
	chatWebsocketCtrl.m.HandleMessage(chatWebsocketCtrl.HandleBroadcastingMessage)
	return &chatWebsocketCtrl
}

func (controller *ChatWebsktCtrlImpl) ServeWebsocketRequests(c *gin.Context) {
	err := controller.m.HandleRequest(c.Writer, c.Request)
	if err != nil {
		c.Error(err)
	}
}

func (controller *ChatWebsktCtrlImpl) HandleConnect(s *melody.Session) {
	// use s.Request. to get authentication id and then get user and set key, the user id being string key and user object being value 
	uid := authenticationController.AuthorizeWebsocketConnection(s.Request)
	user, err := controller.userService.GetUserByUid(uid)
	if err != nil {
		s.CloseWithMsg([]byte(err.Error()))
	}
	s.Set(strconv.Itoa(user.ID), s)
}

func (controller *ChatWebsktCtrlImpl) HandleBroadcastingMessage(s *melody.Session, msg []byte) {
	var messageRequest messages.MessageRequest
	err := json.Unmarshal(msg, &messageRequest)
	if err != nil {
		s.Close()
	}
	message, err := controller.messageService.CreateMessage(messageRequest)
	if err != nil {
		s.CloseWithMsg([]byte(err.Error()))
	}
	messageResponseChan := make(chan messages.MessageResponse)
    errMessageResponse := make(chan error)
	go func (message *ent.Message) {
		messageResponseConcurrentFunc, err := messages.EntToResponse(message)
		messageResponseChan <- messageResponseConcurrentFunc
		errMessageResponse <- err
	}(message)
	users, err := controller.userService.FindUsersByChatId(messageRequest.ChatId)
	if err != nil {
		s.Close()
	}
	userSessions := make([]*melody.Session, 0)
	for _, user := range users {
		sess, ok := s.Get(strconv.Itoa(user.ID))
		if ok {
			session, ok := sess.(*melody.Session)
			if ok {
				userSessions = append(userSessions, session)
			}
		}
	}
	messageResponse, err := <-messageResponseChan, <-errMessageResponse
	if err != nil {
		s.Close()
	}
	messageRespnoseJSON, err := json.Marshal(messageResponse)
	if err != nil {
		s.Close()
	}
	err = controller.m.BroadcastMultiple(messageRespnoseJSON, userSessions)
	if err != nil {
		s.CloseWithMsg([]byte(err.Error()))
	}
}
