package main

import (
	"flag"
	"net"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/qrug/go-exercises/grpc-exercises/ex3-streamserver/proto"
	"google.golang.org/grpc"
)

// 服务端流式RPC：客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。

// StreamService xx
type StreamService struct {
}

// DownloadPackage xx
func (s *StreamService) DownloadPackage(req *proto.SimpleRequest, sr proto.StreamService_DownloadPackageServer) error {
	glog.Infof("recieve request : %v", req.Data)
	for i := 0; i < 5; i++ {
		err := sr.Send(&proto.SimpleResponse{
			Code:   200,
			Result: strconv.Itoa(i),
		})
		if err != nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		glog.Errorf("net listen fail : %v ", err)
	}
	glog.Infof("net %v listening...", ":8001")
	grpcServer := grpc.NewServer()
	proto.RegisterStreamServiceServer(grpcServer, &StreamService{})
	err2 := grpcServer.Serve(listener)
	if err2 != nil {
		glog.Errorf("grpc server start fail : %v", err)
	}
	glog.Infof("server running ...")
}
