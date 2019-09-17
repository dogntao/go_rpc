package main

import (
	"context"
	gs "go_rpc/go_rpc"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct{}

func (s *server) HelloWorld(ctx context.Context, rq *gs.HelloRequest) (rp *gs.HelloRes, err error) {
	log.Printf("Receive:%v", rq.GetReq())
	rp = &gs.HelloRes{Resp: rq.GetReq()}
	err = nil
	return
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("%s", err)
	}
	// 创建grpc服务器实例
	s := grpc.NewServer()
	// 注册服务
	gs.RegisterGreeterServer(s, &server{})
	// 阻塞等待
	if err := s.Serve(lis); err != nil {
		log.Fatalf("%v", err)
	}
}
