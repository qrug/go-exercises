package main

import (
	"context"
	"io"

	"github.com/golang/glog"
	"github.com/qrug/go-exercises/grpc-exercises/ex3-streamserver/proto"
	"google.golang.org/grpc"
)

// Address xx
const Address = ":8001"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		glog.Errorf("connect to server %v fail!", Address)
	}

	grpcClient := proto.NewStreamServiceClient(conn)

	req := proto.SimpleRequest{
		Data: "give me a serious of numbers,please",
	}
	stream, err := grpcClient.DownloadPackage(context.Background(), &req)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			glog.Errorf("stream recieve dataa fail : %v", err)
		}
		glog.Infof("recieve [%v] from stream", res)
	}
	glog.Infof("request finish!")
}
