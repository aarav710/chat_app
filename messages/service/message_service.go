package service

import (
	"chatapp/backend/ent"
	"chatapp/backend/messages/repo"
	"chatapp/backend/messages"
	chatRepo "chatapp/backend/chats/repo"
	userRepo "chatapp/backend/users/repo"
	"chatapp/backend/errors"
)

type MessageService interface {
    FindMessagesByChatId(uid string, chatId, cursor int) ([]*ent.Message, error)
	CreateMessage(messageRequest messages.MessageRequest, chatId int, user *ent.User) (*ent.Message, error)
}

type MessageServiceImpl struct {
	messageRepo repo.MessageRepo
	chatRepo chatRepo.ChatRepo
	userRepo userRepo.UserRepo
}

func NewMessageService(messageRepo repo.MessageRepo, chatRepo chatRepo.ChatRepo, userRepo userRepo.UserRepo) MessageService {
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

func (service *MessageServiceImpl) CreateMessage(messageRequest messages.MessageRequest, chatId int, user *ent.User) (*ent.Message, error) {
	isUserInChat, err := service.userRepo.IsUserInChat(user.ID, chatId)
	if err != nil {
		return nil, err
	}
	if !isUserInChat {
		return nil, errors.UnauthorizedError
	}
	message, err := service.messageRepo.CreateMessage(messageRequest, user.ID, chatId)
	return message, err
}