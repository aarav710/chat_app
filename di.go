//+build wireinject

package main

import (
	chatsDI "chatapp/backend/chats/di"
	userAuth "chatapp/backend/auth"
	"chatapp/backend/di"
	"chatapp/backend/ent"
	usersDI "chatapp/backend/users/di"
	"context"
	loginDI "chatapp/backend/login/di"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeDI(context context.Context, router *gin.Engine, db *ent.Client, auth *auth.Client) di.Controllers{
    wire.Build(usersDI.UsersSet, chatsDI.ChatsSet, loginDI.LoginSet, userAuth.NewAuthService, di.NewControllers)
    return di.Controllers{}
}