package chats

import (
	"chatapp/backend/ent"
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

func EntToChatResponse(chat *ent.Chat) ChatResponse {
  chatResponse := ChatResponse{}
  chatResponse.ID = chat.ID
  chatResponse.Title = chat.Title
  chatResponse.DisplayPictureUrl = chat.DisplayPictureURL
  chatResponse.Description = chat.Description
  chatResponse.CreatedAt = chat.CreatedAt
  return chatResponse
}

