package hub

import (
	"chatapp/backend/ent"
	userService "chatapp/backend/users/service"
)

type HubImpl struct {
	Broadcast chan MessageBroadcast
    Register chan UserBroadcast
	Unregister chan *Client
	userService userService.UserService
	Clients map[*Client]*ent.User
}

type MessageBroadcast struct {
	message *ent.Message
	clients []*Client
}

type UserBroadcast struct {
	Client *Client
	User *ent.User
}

type Hub interface {
	BroadcastMessage(message *ent.Message, chatId int, user *ent.User) (*ent.Message, error)
	UserJoin(user *ent.User) error
	UserUnregister(uid string) error
}

func NewHub(userService userService.UserService) Hub {
	hub := HubImpl{ 
		userService: userService,
		Broadcast: make(chan MessageBroadcast), 
		Register: make(chan UserBroadcast), 
		Unregister: make(chan *Client), 
		Clients: make(map[*Client]*ent.User),
	}
	return &hub
}

func (hub *HubImpl) BroadcastMessage(message *ent.Message, chatId int, user *ent.User) (*ent.Message, error) {
	users, err := hub.userService.FindUsersByChatId(chatId)
	if err != nil {
		return nil, err
	}
	var clients []*Client
	for _, user := range users {
		for client, online_user := range hub.Clients {
			if user.ID == online_user.ID {
				clients = append(clients, client)
			}
		}
	}
	messageBroadcast := MessageBroadcast{clients: clients, message: message}
	hub.Broadcast <- messageBroadcast
	return nil, nil
}

func (hub *HubImpl) UserJoin(user *ent.User) error {
	return nil
}

func (hub *HubImpl) UserUnregister(uid string) error {
	return nil
}