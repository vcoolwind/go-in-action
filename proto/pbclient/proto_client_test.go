package pbclient

import (
	"context"
	"fmt"
	"go-in-action/proto/pb"
	"google.golang.org/grpc"
	"testing"
)

func TestClient(t *testing.T) {
	serviceAddress := "127.0.0.1:1234"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()

	userClient := pb.NewUserServiceClient(conn)

	userReq := &pb.LoginRequest{Username: "123", Password: "1qaz"}
	reply, _ := userClient.CheckPassword(context.Background(), userReq)
	fmt.Printf("UserService CheckPassword : %s, %s \n", reply.Ret,reply.Err)

	userReq2 := &pb.LoginRequest{Username: "admin", Password: "admin"}
	reply2, _ := userClient.CheckPassword(context.Background(), userReq2)
	fmt.Printf("UserService CheckPassword : %s, %s \n", reply2.Ret,reply2.Err)

}
