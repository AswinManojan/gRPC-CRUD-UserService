package interfaces

import model "github.com/grpc/gobus/user/pkg/Model"

type UserRepoInter interface {
	CreateUser(user *model.User) (*model.User, error)
	EditUser(user *model.User) (*model.User, error)
	DeleteUserById(id uint) (*model.User, error)
	FindUserById(id uint) (*model.User, error)
	FindUserByName(name string) (*model.User, error)
}
