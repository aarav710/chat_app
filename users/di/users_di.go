package di

import (
	"chatapp/backend/users/controller"
	"chatapp/backend/users/repo"
	"chatapp/backend/users/service"

	"github.com/google/wire"
)

var UsersSet = wire.NewSet(controller.NewUserController, 
	wire.Bind(new(controller.UserController), 
	          new(*controller.UserControllerImpl)), 
	service.NewUserService, 
	wire.Bind(new(service.UserService), 
	          new(*service.UserServiceImpl)),
	repo.NewUserRepo,
    wire.Bind(new(repo.UserRepo), 
	          new(*repo.UserRepoImpl)))
