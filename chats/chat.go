package chats

import (
	"chatapp/backend/ent"
	"chatapp/backend/errors"
	messageMappings "chatapp/backend/messages"
	"time"
)

type ChatRequest struct {
  ChatBase
}

type ChatBase struct {
  Title string `json:"title" validate:"max=100"`
  DisplayPictureUrl string `json:"display_picture_url"`
  Description string `json:"description" validate:"max=255"`
}

type ChatResponse struct {
  ChatBase
  CreatedAt time.Time `json:"created_at"`
  ID int `json:"id"`
}

type ChatDetailResponse struct {
  ChatResponse
  LatestMessage *messageMappings.MessageResponse `json:"latest_message"`
}

func EntToChatResponse(chat *ent.Chat) ChatResponse {
  chatResponse := ChatResponse{}
  chatResponse.ID = chat.ID
  chatResponse.Title = chat.Title
  chatResponse.DisplayPictureUrl = chat.DisplayPictureURL
  chatResponse.Description = chat.Description
  chatResponse.CreatedAt = chat.CreatedAt
  return chatResponse
}

func EntToChatDetailResponse(chat *ent.Chat) (ChatDetailResponse, error) {
  chatResponse := ChatDetailResponse{}
  chatResponse.ID = chat.ID
  chatResponse.Title = chat.Title
  chatResponse.DisplayPictureUrl = chat.DisplayPictureURL
  chatResponse.Description = chat.Description
  chatResponse.CreatedAt = chat.CreatedAt
  if chat.Edges.Messages == nil {
    return chatResponse, errors.InternalServerError
  } else if len(chat.Edges.Messages) == 0 {
    chatResponse.LatestMessage = nil
  } else {
    latestMessageResponse, err := messageMappings.EntToResponse(chat.Edges.Messages[0])
    if err != nil {
      return chatResponse, err
    }
    chatResponse.LatestMessage = &latestMessageResponse
  }
  return chatResponse, nil
}