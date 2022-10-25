package messages

import (
	"chatapp/backend/ent"
    "chatapp/backend/messages/service"
	"chatapp/backend/messages"
)

type HubImpl struct {
	Broadcast chan messages.MessageRequest
    Register chan *ent.User
	Unregister chan int
    messageService service.MessageService
}

type Hub interface {
	BroadcastMessage(message messages.MessageRequest, chatId int, user *ent.User) (*ent.Message, error)
	UserJoin(user *ent.User) error
	UserUnregister(uid string) error
}

func NewHub(messageService service.MessageService) Hub {
	hub := HubImpl{messageService: messageService}
	return &hub
}

func (hub *HubImpl) BroadcastMessage(messageRequest messages.MessageRequest, chatId int, user *ent.User) (*ent.Message, error) {
	message, err := hub.messageService.CreateMessage(messageRequest, chatId, user)
	if err != nil {
		return nil, err
	}
	
	return message, nil
}

func (hub *HubImpl) UserJoin(user *ent.User) error {
	return nil
}

func (hub *HubImpl) UserUnregister(uid string) error {
	return nil
}