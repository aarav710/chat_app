package repo

import (
	"chatapp/backend/ent"
	"chatapp/backend/ent/chat"
	"chatapp/backend/ent/login"
	"chatapp/backend/ent/user"
	"context"
	userMappings "chatapp/backend/users"
)


type UserRepo interface {
	GetUserById(id int) (*ent.User, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetUsersContainingUsername(username string) ([]*ent.User, error)
	GetUserByUid(uid string) (*ent.User, error)
	FindUsersByChatId(chatId int) ([]*ent.User, error)
	UpdateUser(userRequest userMappings.UserRequest, userId int) (*ent.User, error)
	CreateUser(userRequest userMappings.UserRequest, login *ent.Login) (*ent.User, error)
	DeleteUser(userId int) error
	AddUserToChat(userId, chatId int) (*ent.User, error)
	RemoveUserFromChat(userId, chatId int) error
	IsUserInChat(userId, chatId int) (bool, error)
}

type UserRepoImpl struct {
	ctx context.Context
	db *ent.Client
}

func NewUserRepo(ctx context.Context, db *ent.Client) UserRepo {
	return &UserRepoImpl{ctx: ctx, db: db}
}

func (repo *UserRepoImpl) GetUserById(id int) (*ent.User, error) {
  user, err := repo.db.User.Query().Where(user.ID(id)).WithLogin().Only(repo.ctx)
	return user, err
}

func (repo *UserRepoImpl) GetUserByEmail(email string) (*ent.User, error) {
	user, err := repo.db.Login.Query().Where(login.Email(email)).QueryUser().WithLogin().Only(repo.ctx)
	return user, err
}

func (repo *UserRepoImpl) GetUsersContainingUsername(username string) ([]*ent.User, error) {
	users, err := repo.db.Login.Query().Where(login.UsernameContains(username)).QueryUser().WithLogin().All(repo.ctx)
	return users, err
}

func (repo *UserRepoImpl) GetUserByUid(uid string) (*ent.User, error) {
	user, err := repo.db.Login.Query().Where(login.UID(uid)).QueryUser().WithLogin().Only(repo.ctx)
	return user, err
}

func (repo *UserRepoImpl) FindUsersByChatId(chatId int) ([]*ent.User, error) {
	users, err := repo.db.Chat.Query().Where(chat.ID(chatId)).QueryUsers().WithLogin().All(repo.ctx)
	return users, err
}

func (repo *UserRepoImpl) UpdateUser(userRequest userMappings.UserRequest, userId int) (*ent.User, error) {
    user, err := repo.db.User.UpdateOneID(userId).SetBio(*userRequest.Bio).Save(repo.ctx)
	user.Edges.Login = user.QueryLogin().OnlyX(repo.ctx)
	return user, err
}

func (repo *UserRepoImpl) CreateUser(userRequest userMappings.UserRequest, login *ent.Login) (*ent.User, error) {
	user, err := repo.db.User.Create().SetBio(*userRequest.Bio).SetDisplayPictureURL(userRequest.DisplayPictureUrl).SetLogin(login).Save(repo.ctx)
	return user, err
}

func (repo *UserRepoImpl) DeleteUser(userId int) error {
	err := repo.db.User.DeleteOneID(userId).Exec(repo.ctx)
	return err
}

func (repo *UserRepoImpl) AddUserToChat(userId, chatId int) (*ent.User, error) {
	user, err := repo.db.User.UpdateOneID(userId).AddChatIDs(chatId).Save(repo.ctx)
	user.Edges.Login = user.QueryLogin().OnlyX(repo.ctx)
    return user, err
}

func (repo *UserRepoImpl) RemoveUserFromChat(userId, chatId int) error {
	_, err := repo.db.User.UpdateOneID(userId).RemoveChatIDs(chatId).Save(repo.ctx)
	return err
}

func (repo *UserRepoImpl) IsUserInChat(userId, chatId int) (bool, error) {
	isUserInChat, err := repo.db.Chat.Query().
	                                    Where(chat.And(
											chat.ID(chatId)), 
	                                        chat.HasUsersWith(user.ID(userId))).
	                                    Exist(repo.ctx)
	return isUserInChat, err
}