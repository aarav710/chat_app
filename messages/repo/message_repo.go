package repo

import (
	"chatapp/backend/ent"
	"chatapp/backend/ent/chat"
	messageMappings "chatapp/backend/messages"
	"context"
)

type MessageRepoImpl struct {
  ctx context.Context
  db *ent.Client
}

type MessageRepo interface {
	FindMessagesByChatId(offset, limit, chatId int) ([]*ent.Message, error)
    CountMessagesByChatId(chatId int) (int, error)
	CreateMessage(messageRequest messageMappings.MessageRequest, userId, chatId int) (*ent.Message, error)
}

func NewMessageRepo(ctx context.Context, db *ent.Client) MessageRepo {
	messageRepo := MessageRepoImpl{ctx: ctx, db: db}
	return &messageRepo
}

func (repo *MessageRepoImpl) FindMessagesByChatId(offset, limit, chatId int) ([]*ent.Message, error) {
	messages, err := repo.db.Chat.Query().Where(chat.ID(chatId)).
	                                    QueryMessages().
	                                    Order(ent.Asc(chat.FieldCreatedAt)).
	                                    Offset(offset).
	                                    Limit(limit).All(repo.ctx)
	return messages, err
}

func (repo *MessageRepoImpl) CountMessagesByChatId(chatId int) (int, error) {
	messagesCount, err := repo.db.Chat.Query().Where(chat.ID(chatId)).QueryMessages().Count(repo.ctx)
	return messagesCount, err
}

func (repo *MessageRepoImpl) CreateMessage(messageRequest messageMappings.MessageRequest, userId, chatId int) (*ent.Message, error) {
	message, err := repo.db.Message.Create().
	                                        SetChatID(chatId).
	                                        SetUserID(userId).
	                                        SetText(messageRequest.Text).
	                                        Save(repo.ctx)
	return message, err
}