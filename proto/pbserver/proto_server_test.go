package pbclient

import (
	"context"
	"flag"
	"go-in-action/proto/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

type UserService struct{}

func (s *UserService) CheckPassword(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println(req.Username,req.Password)
	if req.Username == "admin" && req.Password == "admin" {
		response := pb.LoginResponse{Ret: "success",Err:""}
		return &response, nil
	}

	response := pb.LoginResponse{Ret: "fail",Err:"name or password error"}
	return &response, nil
}

func TestServer(t *testing.T) {
	flag.Parse()

	lis, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userService := UserService{}
	pb.RegisterUserServiceServer(grpcServer, &userService)
	log.Println("server will be start ...")
	grpcServer.Serve(lis)
}
