package di

import (
	config "github.com/grpc/gobus/user/Config"
	db "github.com/grpc/gobus/user/pkg/DB"
	handler "github.com/grpc/gobus/user/pkg/Handler"
	repo "github.com/grpc/gobus/user/pkg/Repo"
	server "github.com/grpc/gobus/user/pkg/Server"
	service "github.com/grpc/gobus/user/pkg/Service"
)

func Init() {
	config := config.LoadConfig()
	Db := db.Connect_DB(config)
	userRepo := repo.NewUserRepoImpl(Db)
	userService := service.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService)
	server.NewGRPCServer(config, userHandler)
}
