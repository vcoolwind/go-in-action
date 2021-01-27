package pbclient

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-in-action/proto/pb"
	"google.golang.org/grpc"
	"net"
	"testing"
)

// 声明 Transfer 结构体
type Transfer struct {
	Conn          net.Conn       // 连接
	Buf           [1024 * 2]byte // 传输时，使用的缓冲
}


// 获取并解析服务器的消息
func (transfer *Transfer) ReadResponse(response *pb.LoginResponse) (err error) {
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		return
	}

	// 根据 buf[:4] 转成一个 uint32 类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[:4])
	//根据pkglen 读取消息内容
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	if err = proto.Unmarshal(transfer.Buf[:pkgLen], response); err != nil {
		return
	}
	return
}

// 发送消息到服务器
func (transfer *Transfer) SendMsg(action *pb.LoginRequest) (err error) {
	var (
		sendBytes []byte
		readLen   int
	)
	//sendBytes, ints := action.Descriptor()
	if sendBytes, err = proto.Marshal(action); err != nil {
		return
	}

	pkgLen := uint32(len(sendBytes))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4],pkgLen)

	if readLen, err = transfer.Conn.Write(buf[:4]); readLen != 4 && err != nil {
		if readLen == 0 {
			return errors.New("发送数据长度发生异常，长度为0")
		}
		return
	}
	// 发送消息
	if readLen, err = transfer.Conn.Write(sendBytes); err != nil {
		if readLen == 0 {
			return errors.New("检查到服务器关闭，客户端也关闭")
		}
		return
	}
	return
}

func TestClientJava(t *testing.T) {
	serviceAddress := "127.0.0.1:2234"
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
