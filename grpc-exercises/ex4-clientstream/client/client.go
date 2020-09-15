package main

import (
	"context"
	"fmt"
	"io"

	"github.com/labstack/gommon/log"
	"github.com/qrug/go-exercises/grpc-exercises/ex4-clientstream/proto"
	"google.golang.org/grpc"
)

var streamClient proto.ClientStreamClient

func main() {
	conn, err := grpc.Dial(":8004",grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("connect to server fail : %v", err)
	}
	log.Print("connect to :8004 success	!")

	streamClient = proto.NewClientStreamClient(conn)
	routeHello()
}

func routeHello()  {
	stream,err := streamClient.HelloServer(context.Background())
	if err!=nil {
		log.Printf("fail to send messge , err : %v", err)
	}
	for i := 0; i < 5; i++ {
		err := stream.Send(&proto.HelloRequest{RequestId:2345,Data:fmt.Sprintf("hello %v",i)})
		// 发送也要检测EOF，当服务端在消息没接收完前主动调用SendAndClose()关闭stream，此时客户端还执行Send()，则会返回EOF错误，所以这里需要加上io.EOF判断
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("send stream request fail : %v",err)
		}
	}
	response ,err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("recieve response fail : %v",err)
	}
	log.Printf("recieve response : [%v]",response)
}