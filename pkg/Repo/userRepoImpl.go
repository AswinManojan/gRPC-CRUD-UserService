package repo

import (
	"fmt"
	"log"

	model "github.com/grpc/gobus/user/pkg/Model"
	interfaces "github.com/grpc/gobus/user/pkg/Repo/Interfaces"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

// CreateUser implements interfaces.UserRepoInter.
func (ur *UserRepo) CreateUser(user *model.User) (*model.User, error) {
	result := ur.DB.Create(user)
	if result.Error != nil {
		log.Print("Error creating user - UserRepo")
		return nil, result.Error
	}
	return user, nil
}

// DeleteUserById implements interfaces.UserRepoInter.
func (ur *UserRepo) DeleteUserById(id uint) (*model.User, error) {
	var usr *model.User
	result := ur.DB.Where("id=?", id).Delete(&usr)
	if result.Error != nil {
		log.Print("Error Deleting the user - UserRepo")
		return nil, result.Error
	}
	return usr, nil
}

// EditUser implements interfaces.UserRepoInter.
func (ur *UserRepo) EditUser(user *model.User) (*model.User, error) {
	ID := user.ID
	fmt.Println(ID)
	fmt.Println(user)
	var usr model.User
	result := ur.DB.Where("username=?", user.Username).First(&usr)
	if result.Error != nil {
		log.Print("Error finding the user")
		return nil, result.Error
	}
	usr.Password = user.Password
	usr.Email = user.Email
	if err := ur.DB.Save(&usr).Error; err != nil {
		log.Print("Error saving/updating the user")
		return nil, err
	}
	return &usr, nil
}

// FindUserById implements interfaces.UserRepoInter.
func (ur *UserRepo) FindUserById(id uint) (*model.User, error) {
	var user model.User
	result := ur.DB.Where("id=?", id).Find(&user)
	if result.Error != nil {
		log.Print("Error finding the user by id")
		return nil, result.Error
	}
	return &user, nil
}

// FindUserByName implements interfaces.UserRepoInter.
func (ur *UserRepo) FindUserByName(name string) (*model.User, error) {
	var user model.User
	fmt.Println(user)
	result := ur.DB.Where("username=?", name).Find(&user)
	if result.Error != nil {
		log.Print("Error finding the user by name")
		return nil, result.Error
	}
	fmt.Println(user)
	return &user, nil
}

func NewUserRepoImpl(db *gorm.DB) interfaces.UserRepoInter {
	return &UserRepo{
		DB: db,
	}
}
