package users

import (
	"chatapp/backend/ent"
	"chatapp/backend/errors"
	"fmt"
	"time"
)

type UserRequest struct {
  UserBase
}

type UserBase struct {
  Bio *string `json:"bio" validate:"max=255"`
  DisplayPictureUrl string `json:"display_picture_url"`
}

type UserResponse struct {
  ID int `json:"id"`
  Username string `json:"username"`
	UserBase
  CreatedAt time.Time `json:"created_at"`
  Status string `json:"status"`
}

func RequestToEnt(request UserRequest) *ent.User {
  model := &ent.User{}
  model.Bio = *request.Bio
  model.DisplayPictureURL = request.DisplayPictureUrl
  return model
}


func EntToResponse(entity *ent.User) (UserResponse, error) {
  response := UserResponse{}
  response.Bio = &entity.Bio
  response.ID = entity.ID
  response.DisplayPictureUrl = entity.DisplayPictureURL
  if entity.Edges.Login == nil {
    fmt.Println("whoops")
    return response, errors.InternalServerError
  }
  response.Username = entity.Edges.Login.Username
  response.CreatedAt = entity.Edges.Login.CreatedAt
  response.Status = string(entity.Edges.Login.Status)
  return response, nil
}