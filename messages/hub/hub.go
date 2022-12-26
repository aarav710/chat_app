package hub

import (
	"chatapp/backend/ent"
	userService "chatapp/backend/users/service"
	messageService "chatapp/backend/messages/service"
	messageMappings "chatapp/backend/messages"
)

type HubImpl struct {
	Broadcast chan messageMappings.MessageRequest
    Register chan UserBroadcast
	Unregister chan *Client
	userService userService.UserService
	messageService messageService.MessageService
	Clients map[*Client]*ent.User
}

type SendMessage struct {
	message messageMappings.MessageResponse
	clients []*Client
}

type UserBroadcast struct {
	Client *Client
	user *ent.User
}

type Hub interface {
	BroadcastMessage(messageRequest messageMappings.MessageRequest) (messageMappings.MessageResponse, error)
	UserJoin(UserBroadcast UserBroadcast) error
	Start()
	UserLeave(client *Client) error 
}

func NewHub(userService userService.UserService, messageService messageService.MessageService) Hub {
	hub := HubImpl{ 
		userService: userService,
		messageService: messageService,
		Broadcast: make(chan messageMappings.MessageRequest), 
		Register: make(chan UserBroadcast), 
		Unregister: make(chan *Client), 
		Clients: make(map[*Client]*ent.User),
	}
	return &hub
}

func (hub *HubImpl) Start() {
	for {
		select {
		case message := <- hub.Broadcast:
			_, err := hub.BroadcastMessage(message)
			if err != nil {
				panic("")
			}
		case userMessReq := <- hub.Register:
			hub.Clients[userMessReq.Client] = userMessReq.user
		case client := <- hub.Unregister:
			if _, ok := hub.Clients[client]; ok {
				delete(hub.Clients, client)
				close(client.send)
			}
		}

	}
}

func (hub *HubImpl) BroadcastMessage(messageRequest messageMappings.MessageRequest) (messageMappings.MessageResponse, error) {
	users, err := hub.userService.FindUsersByChatId(messageRequest.ChatId)
	if err != nil {
		return messageMappings.MessageResponse{}, err
	}
	message, err := hub.messageService.CreateMessage(messageRequest)
	if err != nil {
		return messageMappings.MessageResponse{}, err
	}
	messageResponse, err := messageMappings.EntToResponse(message)
	if err != nil {
		return messageMappings.MessageResponse{}, err
	}
	for _, user := range users {
		for client, online_user := range hub.Clients {
			if user.ID == online_user.ID && user.ID != messageRequest.SenderId {
				client.send <- messageResponse
			}
		}
	}
	return messageResponse, nil
}

func (hub *HubImpl) UserJoin(userBroadcast UserBroadcast) error {
	hub.Register <- userBroadcast
	return nil
}

func (hub *HubImpl) UserLeave(client *Client) error {
	hub.Unregister <- client
	return nil
}