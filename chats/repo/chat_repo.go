package repo

import (
	chatMappings "chatapp/backend/chats"
	"chatapp/backend/ent"
	"chatapp/backend/ent/chat"
	"chatapp/backend/ent/message"
	"chatapp/backend/ent/user"
	"context"
)

type ChatRepoImpl struct {
  db *ent.Client
  ctx context.Context
}

type ChatRepo interface {
    FindChatsByUserId(userId int) ([]*ent.Chat, error)
	FindChatById(chatId int) (*ent.Chat, error)
	CreateChat(chatRequest chatMappings.ChatRequest, userId int) (*ent.Chat, error)
	UpdateChat(chatRequest chatMappings.ChatRequest, chatId int) (*ent.Chat, error)
	DeleteChat(chatId int) error
}

func NewChatRepo(db *ent.Client, ctx context.Context) *ChatRepoImpl {
	chatRepo := ChatRepoImpl{db: db, ctx: ctx}
    return &chatRepo
}

func (repo *ChatRepoImpl) FindChatsByUserId(userId int) ([]*ent.Chat, error) {
	chats, err := repo.db.Message.Query().
	                                    Order(ent.Desc(message.FieldCreatedAt)).
	                                    QueryChat().
	                                    Where(chat.HasUsersWith(user.ID(userId))).
	                                    All(repo.ctx)
	return chats, err
}

func (repo *ChatRepoImpl) FindChatById(chatId int) (*ent.Chat, error) {
	chat, err := repo.db.Chat.Query().Where(chat.ID(chatId)).Only(repo.ctx)
	return chat, err
}

func (repo *ChatRepoImpl) CreateChat(chatRequest chatMappings.ChatRequest, userId int) (*ent.Chat, error) {
    chat, err := repo.db.Chat.Create().
	                                AddUserIDs(userId).
	                                SetDescription(chatRequest.Description).
	                                SetDisplayPictureURL(chatRequest.DisplayPictureUrl).
	                                SetTitle(chatRequest.Title).
									Save(repo.ctx)
	return chat, err
}

func (repo *ChatRepoImpl) UpdateChat(chatRequest chatMappings.ChatRequest, chatId int) (*ent.Chat, error) {
	updatedChat, err := repo.db.Chat.UpdateOneID(chatId).
	                                            SetDescription(chatRequest.Description).
	                                            SetTitle(chatRequest.Title).
												SetDisplayPictureURL(chatRequest.DisplayPictureUrl).
	                                            Save(repo.ctx)
	return updatedChat, err
}

func (repo *ChatRepoImpl) DeleteChat(chatId int) error {
	err := repo.db.Chat.DeleteOneID(chatId).Exec(repo.ctx)
	return err
}