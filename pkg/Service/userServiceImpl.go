package service

import (
	"errors"
	"log"

	model "github.com/grpc/gobus/user/pkg/Model"
	Repointerface "github.com/grpc/gobus/user/pkg/Repo/Interfaces"
	interfaces "github.com/grpc/gobus/user/pkg/Service/Interfaces"
)

type UserServiceImpl struct {
	repo Repointerface.UserRepoInter
}

// AddUser implements interfaces.UserServiceInterface.
func (us *UserServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	result, err := us.repo.CreateUser(user)
	if err != nil {
		log.Print("Error while creating user in Service")
		return nil, err
	}
	return result, nil
}

// DeleteUserByID implements interfaces.UserServiceInterface.
func (us *UserServiceImpl) DeleteUserByID(id uint) (*model.User, error) {
	result, err := us.repo.DeleteUserById(id)
	if err != nil {
		log.Print("Error while deleting user in Service")
		return nil, err
	}
	return result, nil
}

// EditUserInfo implements interfaces.UserServiceInterface.
func (us *UserServiceImpl) EditUserInfo(user *model.User) (*model.User, error) {
	result, err := us.repo.EditUser(user)
	if err != nil {
		log.Print("Error while editing user in Service")
		return nil, err
	}
	return result, nil
}

// UserLogin implements interfaces.UserServiceInterface.
func (us *UserServiceImpl) UserLogin(userLogin *model.UserLogin) (*model.User, error) {
	user, err := us.repo.FindUserByName(userLogin.Username)
	if err != nil {
		log.Print("Error while finding user in Service")
		return nil, err
	}
	if user.Password != userLogin.Password {
		log.Print("Invalid login credentials")
		return nil, errors.New("password Mismatch")
	}
	return user, nil
}

// ViewUserByID implements interfaces.UserServiceInterface.
func (us *UserServiceImpl) ViewUserByID(id uint) (*model.User, error) {
	result, err := us.repo.FindUserById(id)
	if err != nil {
		log.Print("Error while finding user in Service")
		return nil, err
	}
	return result, nil
}

func NewUserServiceImpl(REPO Repointerface.UserRepoInter) interfaces.UserServiceInterface {
	return &UserServiceImpl{
		repo: REPO,
	}
}
