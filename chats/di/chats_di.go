package di

import (
	"chatapp/backend/chats/controller"
	"chatapp/backend/chats/repo"
	"chatapp/backend/chats/service"

	"github.com/google/wire"
)

var ChatsSet = wire.NewSet(controller.NewChatController, 
	wire.Bind(new(controller.ChatController), 
	          new(*controller.ChatControllerImpl)), 
	service.NewChatService, 
	wire.Bind(new(service.ChatService), 
	          new(*service.ChatServiceImpl)),
	repo.NewChatRepo,
    wire.Bind(new(repo.ChatRepo), 
	          new(*repo.ChatRepoImpl)))