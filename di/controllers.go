package di

import (
	chatController "chatapp/backend/chats/controller"
	userController "chatapp/backend/users/controller"
)

type Controllers struct {
	chatController chatController.ChatController
	userController userController.UserController
}

func NewControllers(chatController chatController.ChatController, userController userController.UserController) Controllers {
	return Controllers{chatController: chatController, userController: userController}
}