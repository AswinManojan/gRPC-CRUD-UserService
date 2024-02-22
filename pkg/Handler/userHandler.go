package handler

import (
	"context"
	"fmt"
	"log"

	model "github.com/grpc/gobus/user/pkg/Model"
	pb "github.com/grpc/gobus/user/pkg/PB"
	interfaces "github.com/grpc/gobus/user/pkg/Service/Interfaces"
)

type UserHandler struct {
	service interfaces.UserServiceInterface
	pb.UserServicesServer
}

func (uh *UserHandler) UserLogin(ctx context.Context, pbl *pb.LoginRequest) (*pb.Result, error) {
	userLogin := &model.UserLogin{
		Username: pbl.Username,
		Password: pbl.Password,
	}
	result, err := uh.service.UserLogin(userLogin)
	if err != nil {
		log.Print("Error while user logging in - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "User Login Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - User successfully logged in.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (uh *UserHandler) UserSignUp(ctx context.Context, pbd *pb.UserDataRequest) (*pb.Result, error) {
	userData := &model.User{
		Username: pbd.UserName,
		Password: pbd.Password,
		Email:    pbd.Email,
	}
	result, err := uh.service.CreateUser(userData)
	if err != nil {
		log.Print("Error while user signing up - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "User Signup Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - User successfully Signed in.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (uh *UserHandler) EditUser(ctx context.Context, pbd *pb.UserDataRequest) (*pb.Result, error) {
	userData := &model.User{
		Username: pbd.UserName,
		Password: pbd.Password,
		Email:    pbd.Email,
	}
	result, err := uh.service.EditUserInfo(userData)
	if err != nil {
		log.Print("Error while editing user info - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "Edit User Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - UserInfo successfully edited.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (uh *UserHandler) DeleteUserById(ctx context.Context, pbi *pb.IdRequest) (*pb.Result, error) {
	result, err := uh.service.DeleteUserByID(uint(pbi.Id))
	if err != nil {
		log.Print("Error while Deleting user - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "User deletion Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - UserInfo successfully Deleted.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (uh *UserHandler) FindUserById(ctx context.Context, pbi *pb.IdRequest) (*pb.Result, error) {
	result, err := uh.service.ViewUserByID(uint(pbi.Id))
	if err != nil {
		log.Print("Error while Finding user - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "User find Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("Fetched User data successfully - Username - %s,Email - %s,Password - %s", result.Username, result.Email, result.Password)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (uh *UserHandler) CreateUser(ctx context.Context, pbd *pb.UserDataRequest) (*pb.Result, error) {
	fmt.Println("Here")
	userData := &model.User{
		Username: pbd.UserName,
		Password: pbd.Password,
		Email:    pbd.Email,
	}
	result, err := uh.service.CreateUser(userData)
	if err != nil {
		log.Print("Error while adding user - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "User Addition Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - User successfully added.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}

func NewUserHandler(svc interfaces.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: svc,
	}
}
