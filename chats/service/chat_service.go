package service

import (
	chatMappings "chatapp/backend/chats"
	"chatapp/backend/chats/repo"
	"chatapp/backend/ent"
	"chatapp/backend/errors"
	userRepo "chatapp/backend/users/repo"
)

type ChatServiceImpl struct {
	chatRepo repo.ChatRepo
	userRepo userRepo.UserRepo
}

type ChatService interface {
	FindChatsByUserId(userId int) ([]*ent.Chat, error)
	FindChatById(user *ent.User, chatId int) (*ent.Chat, error)
	CreateChat(user *ent.User, chatRequest chatMappings.ChatRequest) (*ent.Chat, error)
	UpdateChat(chatRequest chatMappings.ChatRequest, chatId int, uid string) (*ent.Chat, error)
	AddUserToChat(user *ent.User, toBeAddedUserId, chatId int) (*ent.User, error)
	RemoveUserFromChat(user *ent.User, toBeRemovedUserId, chatId int) error
	DeleteChat(user *ent.User, chatId int) error
}

func NewChatService(chatRepo repo.ChatRepo, userRepo userRepo.UserRepo) ChatService {
	chatService := ChatServiceImpl{chatRepo: chatRepo, userRepo: userRepo}
	return &chatService
}

func (service *ChatServiceImpl) FindChatsByUserId(userId int) ([]*ent.Chat, error) {
	chats, err := service.chatRepo.FindChatsByUserId(userId)
	return chats, err
}

func (service *ChatServiceImpl) FindChatById(user *ent.User, chatId int) (*ent.Chat, error) {
	isUserInChat, err := service.userRepo.IsUserInChat(user.ID, chatId)
	if err != nil {
		return nil, err
	}
	if !isUserInChat {
		return nil, errors.UnauthorizedError
	}
	chat, err := service.chatRepo.FindChatById(chatId)
	return chat, err
}

func (service *ChatServiceImpl) CreateChat(user *ent.User, chatRequest chatMappings.ChatRequest) (*ent.Chat, error) {
	chat, err := service.chatRepo.CreateChat(chatRequest, user.ID)
	if err != nil {
		return nil, err
	}
	_, addUserError := service.userRepo.AddUserToChat(user.ID, chat.ID)
	if addUserError != nil {
		return nil, err
	}
	return chat, err
}

func (service *ChatServiceImpl) UpdateChat(chatRequest chatMappings.ChatRequest, chatId int, uid string) (*ent.Chat, error) {
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
	chat, err := service.chatRepo.UpdateChat(chatRequest, chatId)
	if err != nil {
		return nil, err
	}
	return chat, err
}

func (service *ChatServiceImpl) AddUserToChat(user *ent.User, toBeAddedUserId, chatId int) (*ent.User, error) {
	isToBeAddedUserInGroup := make(chan bool)
	err1 := make(chan error)

	isUserInGroup := make(chan bool)
	err2 := make(chan error)

	go func() {
		isUserInChat, err := service.userRepo.IsUserInChat(user.ID, chatId)
		isToBeAddedUserInGroup <- isUserInChat
		err1 <- err
	}()

	go func() {
		isUserInChat, err := service.userRepo.IsUserInChat(toBeAddedUserId, chatId)
		isUserInGroup <- isUserInChat
		err2 <- err
	}()
	err := <-err1

	if !<-isToBeAddedUserInGroup {
		return nil, errors.UnauthorizedError
	}
	if err != nil {
		return nil, err
	}

	if <-isToBeAddedUserInGroup {
		return nil, nil // add the already exists dynamic error
	}
	err = <-err2
	if err != nil {
		return nil, err
	}

	addedUser, err := service.userRepo.AddUserToChat(toBeAddedUserId, chatId)
	if err != nil {
		return nil, err
	}
	return addedUser, err
}

func (service *ChatServiceImpl) RemoveUserFromChat(user *ent.User, toBeRemovedUserId, chatId int) error {
	err1 := make(chan error)
	isUserInChat := make(chan bool)
	err2 := make(chan error)
	isUserToBeRemovedInChat := make(chan bool)
	go func() {
        isUserInChatTemp, err := service.userRepo.IsUserInChat(user.ID, chatId)
		isUserInChat <- isUserInChatTemp
		err1 <- err
	}()

	go func() {
        isUserToBeRemovedInChatTemp, err := service.userRepo.IsUserInChat(toBeRemovedUserId, chatId)
		isUserToBeRemovedInChat <- isUserToBeRemovedInChatTemp
		err2 <- err
	}()

    err := <- err1
	
	if err != nil {
		return err
	}
	if <-isUserInChat {
		return errors.UnauthorizedError
	}
	
    err = <- err2

	if err != nil {
		return err
	}
	if !<-isUserToBeRemovedInChat {
		return errors.UnauthorizedError // change to corresponding error type; this is just a placeholder for now
	}
    
	err = service.userRepo.RemoveUserFromChat(toBeRemovedUserId, chatId)
	if err != nil {
		return err
	}
	return nil
}

func (service *ChatServiceImpl) DeleteChat(user *ent.User, chatId int) error {
	isUserInChat, err := service.userRepo.IsUserInChat(user.ID, chatId)
	if err != nil {
		return err
	}
	if !isUserInChat {
		return errors.UnauthorizedError
	}
	err = service.chatRepo.DeleteChat(chatId)
	if err != nil {
		return err
	}
	return nil
}