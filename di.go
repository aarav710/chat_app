//go:build wireinject
// +build wireinject

package main

import (
	userAuth "chatapp/backend/auth"
	chatsDI "chatapp/backend/chats/di"
	"chatapp/backend/ent"
	loginDI "chatapp/backend/login/di"
	messagesDI "chatapp/backend/messages/di"
	"chatapp/backend/messages/hub"
	usersDI "chatapp/backend/users/di"
	"context"
	"chatapp/backend/di"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeDI(context context.Context, router *gin.Engine, db *ent.Client, auth *auth.Client) di.Controllers {
    wire.Build(usersDI.UsersSet, chatsDI.ChatsSet, loginDI.LoginSet, messagesDI.MessagesSet, hub.NewHub, userAuth.NewAuthService, di.NewControllers)
    return di.Controllers{}
}