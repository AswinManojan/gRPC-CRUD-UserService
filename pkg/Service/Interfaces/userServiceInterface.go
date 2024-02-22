package interfaces

import model "github.com/grpc/gobus/user/pkg/Model"

type UserServiceInterface interface {
	UserLogin(userLogin *model.UserLogin) (*model.User, error)
	ViewUserByID(id uint) (*model.User, error)
	EditUserInfo(user *model.User) (*model.User, error)
	DeleteUserByID(id uint) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
}
