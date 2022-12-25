package di

import (
	"chatapp/backend/login/controller"
	"chatapp/backend/login/repo"
	"chatapp/backend/login/service"

	"github.com/google/wire"
)

var LoginSet = wire.NewSet(controller.NewLoginController, 
	wire.Bind(new(controller.LoginController), 
	          new(*controller.LoginControllerImpl)), 
	service.NewLoginService, 
	wire.Bind(new(service.LoginService), 
	          new(*service.LoginServiceImpl)),
	repo.NewLoginRepo,
    wire.Bind(new(repo.LoginRepo), 
	          new(*repo.LoginRepoImpl)))