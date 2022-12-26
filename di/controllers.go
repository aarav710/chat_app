package di

import (
	chatController "chatapp/backend/chats/controller"
	loginController "chatapp/backend/login/controller"
	messageController "chatapp/backend/messages/controller"
	"chatapp/backend/messages/hub"
	userController "chatapp/backend/users/controller"
)

type Controllers struct {
	chatController chatController.ChatController
	userController userController.UserController
	messageController messageController.MessageController
	loginController loginController.LoginController
	Hub hub.Hub
}

func NewControllers(chatController chatController.ChatController, userController userController.UserController, messageController messageController.MessageController, loginController loginController.LoginController, hub hub.Hub) Controllers {
	return Controllers{chatController: chatController, userController: userController, messageController: messageController, loginController: loginController, Hub: hub}
}