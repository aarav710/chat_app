package service

import (
	chatRepo "chatapp/backend/chats/repo"
	"chatapp/backend/ent"
	"chatapp/backend/errors"
	"chatapp/backend/messages"
	"chatapp/backend/messages/repo"
	userRepo "chatapp/backend/users/repo"
)

type MessageService interface {
    FindMessagesByChatId(uid string, chatId, cursor int) ([]*ent.Message, error)
	CreateMessage(messageRequest messages.MessageRequest) (*ent.Message, error)
}

type MessageServiceImpl struct {
	messageRepo repo.MessageRepo
	chatRepo chatRepo.ChatRepo
	userRepo userRepo.UserRepo
}

func NewMessageService(messageRepo repo.MessageRepo, chatRepo chatRepo.ChatRepo, userRepo userRepo.UserRepo) *MessageServiceImpl {
	messageService := MessageServiceImpl{messageRepo: messageRepo, chatRepo: chatRepo, userRepo: userRepo}
	return &messageService
}

// cursor is an integer (taken from the controller's request parameter) that tracks the state of the pages that have been loaded from the infite scroll 
// cursor >= 1
func (service *MessageServiceImpl) FindMessagesByChatId(uid string, chatId, cursor int) ([]*ent.Message, error) {
	user, err := service.userRepo.GetUserByUid(uid)
	if err != nil {
		return nil, err
	}
	isUserInChat, err := service.userRepo.IsUserInChat(user.ID, chatId)
	if err != nil {
		return nil, err
	}
	if !isUserInChat {
		return nil, errors.UnauthorizedError
	}
	messagesCount, err := service.messageRepo.CountMessagesByChatId(chatId)
    if err != nil {
		return nil, err
	}
	offset := messagesCount - (cursor * messages.MESSAGES_LIMIT)
	messages, err := service.messageRepo.FindMessagesByChatId(offset, messages.MESSAGES_LIMIT, chatId)
	return messages, err
}

func (service *MessageServiceImpl) CreateMessage(messageRequest messages.MessageRequest) (*ent.Message, error) {
	user, err := service.userRepo.GetUserById(messageRequest.SenderId)
	if err != nil {
		return nil, err
	}
	isUserInChat, err := service.userRepo.IsUserInChat(messageRequest.SenderId, messageRequest.ChatId)
	if err != nil {
		return nil, err
	}
	if !isUserInChat {
		return nil, errors.UnauthorizedError
	}
	message, err := service.messageRepo.CreateMessage(messageRequest)
	if err != nil {
		return nil, err
	}
	message.Edges.User = user
	return message, nil
}