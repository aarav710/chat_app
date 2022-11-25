package hub

import (
	"chatapp/backend/ent"
	"chatapp/backend/messages"
	"chatapp/backend/messages/service"
)

type HubImpl struct {
	Broadcast chan *ent.Message
    Register chan *ent.User
	Unregister chan int
    messageService service.MessageService
	Clients map[*Client]*ent.User
}


type Hub interface {
	BroadcastMessage(message messages.MessageRequest, chatId int, user *ent.User) (*ent.Message, error)
	UserJoin(user *ent.User) error
	UserUnregister(uid string) error
}

func NewHub(messageService service.MessageService) Hub {
	hub := HubImpl{messageService: messageService, Broadcast: make(chan *ent.Message), Register: make(chan *ent.User), Unregister: make(chan int), Clients: make(map[*Client]*ent.User)}
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