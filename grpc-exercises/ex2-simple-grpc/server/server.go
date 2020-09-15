package main

import (
	"context"
	"log"
	"net"

	pb "github.com/qrug/go-exercises/grpc-exercises/ex2-simple-grpc/proto"
	"google.golang.org/grpc"
)

// gRPC主要有4种请求和响应模式，分别是简单模式(Simple RPC)、服务端流式（Server-side streaming RPC）、
// 客户端流式（Client-side streaming RPC）、和双向流式（Bidirectional streaming RPC）。
// 简单模式(Simple RPC)：客户端发起请求并等待服务端响应。
// 服务端流式（Server-side streaming RPC）：客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。
// 客户端流式（Client-side streaming RPC）：与服务端数据流模式相反，这次是客户端源源不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应。
// 双向流式（Bidirectional streaming RPC）：双方使用读写流去发送一个消息序列，两个流独立操作，双方可以同时发送和同时接收。

// SimpleServer xx
type SimpleServer struct{}

// Route xx
func (s *SimpleServer) Route(context context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatal("net.Listen err :%v", err)
	}
	log.Prinintln(Address + "net listening ...")
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleServer(grpcServer, &SimpleServer{})
	err2 := grpcServer.Serve(listener)
	if err2 != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
