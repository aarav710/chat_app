package users

import "chatapp/backend/ent"

type UserRequest struct {
  UserBase
}

type UserBase struct {
  Bio *string `json:"bio" validate:"max=255"`
}

type UserResponse struct {
  ID int
	UserBase
}

func RequestToEnt(request UserRequest) *ent.User {
  model := &ent.User{}
  model.Bio = *request.Bio
  return model
}


func EntToResponse(entity *ent.User) UserResponse {
  response := UserResponse{}
  response.Bio = &entity.Bio
  response.ID = entity.ID
  return response
}