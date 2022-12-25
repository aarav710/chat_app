package service

import (
	"errors"
	"testing"

	"chatapp/backend/auth"
	"chatapp/backend/ent"
	loginRepo "chatapp/backend/login/repo"
	mockUserRepo "chatapp/backend/mock/users/repo"
	"reflect"

	"github.com/golang/mock/gomock"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	userRepoMock := mockUserRepo.NewMockUserRepo(ctrl)
	loginRepo := loginRepo.LoginRepoImpl{}
	authService := auth.AuthServiceImpl{}
	userId := 5
	t.Run("user does not exist", func(t *testing.T) {
		userRepoMock.EXPECT().GetUserById(userId).Return(nil, errors.New("not found error"))
		userService := NewUserService(userRepoMock, &loginRepo, &authService)
		user, err := userService.GetUserById(userId)
		if user != nil || err == nil {
			t.Errorf("Incorrect response from user service")
		}
	})
	t.Run("user exists", func(t *testing.T) {
		expectUser := &ent.User{
			ID:  5,
			Bio: "hello world!",
		}
		userRepoMock.EXPECT().GetUserById(userId).Return(expectUser, nil)
		userService := NewUserService(userRepoMock, &loginRepo, &authService)
		user, err := userService.GetUserById(userId)
		if !reflect.DeepEqual(*user, *expectUser) {
			t.Errorf("Incorrect response from user service")
		}
		if err != nil {
			t.Errorf("expected no problem from user service.")
		}
	})
}

func TestGetUsersByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	loginRepo := loginRepo.LoginRepoImpl{}
	authService := auth.AuthServiceImpl{}
	defer ctrl.Finish()
	userRepoMock := mockUserRepo.NewMockUserRepo(ctrl)

	username := "john"
	t.Run("users do not exist", func(t *testing.T) {
		var result []*ent.User
		userRepoMock.EXPECT().GetUsersContainingUsername(username).Return(result, nil)
		userService := NewUserService(userRepoMock, &loginRepo, &authService)
		users, err := userService.GetUsersByUsername(username)
		if !reflect.DeepEqual(users, result) || err != nil {
			t.Errorf("Incorrect response from user service")
		}
	})
}



