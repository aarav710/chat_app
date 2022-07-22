package service

import (
	"testing"

  mockUserRepo "chatapp/backend/mock/users/repo"
	"github.com/golang/mock/gomock"
)

func TestGetUserById(t *testing.T) {
  ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	userRepoMock := mockUserRepo.NewMockUserRepo(ctrl)
	userRepoMock.EXPECT().GetUserById(5).Return(nil)
	userService := NewUserService(userRepoMock)
  userService.GetUserById(5)
}