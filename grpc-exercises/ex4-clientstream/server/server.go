package main

import (
	"io"
	"log"
	"net"

	"github.com/qrug/go-exercises/grpc-exercises/ex4-clientstream/proto"
	"google.golang.org/grpc"
)

// SimpleService xx
type SimpleService struct{}

// HelloServer xx
func (s *SimpleService) HelloServer(server proto.ClientStream_HelloServerServer) error {
	for {
		req, err := server.Recv()
		if err == io.EOF {
			return server.SendAndClose(&proto.HelloResponse{Code: 1000, IsSuccess: true, Result: "success"})
		}
		if err != nil {
			return err
		}
		log.Println(req)
	}
}

// Address xx
const Address string = ":8004"

func main() {
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("net listen err : %v", err)
	}
	log.Printf("net listening %v", Address)
	server := grpc.NewServer()
	proto.RegisterClientStreamServer(server, &SimpleService{})
	err2 := server.Serve(listener)
	if err2 != nil {
		log.Printf("server start fail : %v", err)
	}
}
