package main

import (
	"context"

	"github.com/prometheus/common/log"
	pb "github.com/qrug/go-exercises/grpc-exercises/ex2-simple-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal("connect to grpc server fail ! %v", err)
	}

	grpcClient := pb.NewSimpleClient(conn)
	req := &pb.SimpleRequest{
		Data: "grpc",
	}
	res, err := grpcClient.Route(context.Background(), req)
	if err != nil {
		log.Errorf("grpc call fail! %v", err)
	}
	log.Info(res)
}
