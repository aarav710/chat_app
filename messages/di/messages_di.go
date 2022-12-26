package di

import (
	"chatapp/backend/messages/controller"
	"chatapp/backend/messages/repo"
	"chatapp/backend/messages/service"

	"github.com/google/wire"
)

var MessagesSet = wire.NewSet(controller.NewMessageController, 
	wire.Bind(new(controller.MessageController), 
	          new(*controller.MessageControllerImpl)), 
	service.NewMessageService, 
	wire.Bind(new(service.MessageService), 
	          new(*service.MessageServiceImpl)),
	repo.NewMessageRepo,
    wire.Bind(new(repo.MessageRepo), 
	          new(*repo.MessageRepoImpl)))
