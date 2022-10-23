package messages

import (
	"chatapp/backend/ent"
	"chatapp/backend/errors"
	userMappings "chatapp/backend/users"
	"time"
)

var MESSAGES_LIMIT int = 20

type MessageRequest struct {
  MessageBase
}

type MessageBase struct {
  Text string `json:"text"`
}

type MessageResponse struct {
  User userMappings.UserResponse `json:"user"`
  MessageBase
  CreatedAt time.Time `json:"created_at"`
}

func EntToResponse(entity *ent.Message) (MessageResponse, error) {
  messageResponse := MessageResponse{}
  messageResponse.Text = entity.Text
  if entity.Edges.User == nil {
    return messageResponse, errors.InternalServerError
  }
  userResponse, err := userMappings.EntToResponse(entity.Edges.User)
  if err != nil {
    return messageResponse, err
  }
  messageResponse.User = userResponse
  messageResponse.CreatedAt = entity.CreatedAt
  return messageResponse, nil
}