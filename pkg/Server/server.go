package server

import (
	"fmt"
	"log"
	"net"

	config "github.com/grpc/gobus/user/Config"
	handler "github.com/grpc/gobus/user/pkg/Handler"
	pb "github.com/grpc/gobus/user/pkg/PB"
	"google.golang.org/grpc"
)

func NewGRPCServer(cfg *config.Config, handler *handler.UserHandler) error {
	log.Println("Connecting to gRPC Server")
	lis, err := net.Listen("tcp", ":"+cfg.GRPCUSERPORT)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		fmt.Println(err)
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServicesServer(grpcServer, handler)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	fmt.Println("connected")
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
