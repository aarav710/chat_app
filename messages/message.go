package messages

import (
	"chatapp/backend/ent"
	"chatapp/backend/errors"
	userMappings "chatapp/backend/users"
	"time"
)

var MESSAGES_LIMIT int = 20

type MessageRequest struct {
  SenderId int `json:"senderId"`
  MessageBase
}

type MessageBase struct {
  Text string `json:"text" validate:"max=500"`
  ChatId int `json:"chatId"`
}

type MessageResponse struct {
  User userMappings.UserResponse `json:"user"`
  MessageBase
  CreatedAt time.Time `json:"created_at"`
  ID int `json:"id"`
}

func EntToResponse(entity *ent.Message) (MessageResponse, error) {
  messageResponse := MessageResponse{}
  messageResponse.ID = entity.ID
  messageResponse.Text = entity.Text
  if entity.Edges.User == nil {
    return messageResponse, errors.InternalServerError
  }
  userResponse, err := userMappings.EntToResponse(entity.Edges.User)
  if err != nil {
    return messageResponse, err
  }
  if entity.Edges.Chat == nil {
    return messageResponse, errors.InternalServerError
  }
  messageResponse.ChatId = entity.Edges.Chat.ID
  messageResponse.User = userResponse
  messageResponse.CreatedAt = entity.CreatedAt
  return messageResponse, nil
}