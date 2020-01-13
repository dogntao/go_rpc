package main

import (
	"context"
	gs "go_rpc/go_rpc"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct{}

// 实现HelloWorld方法
func (s *server) HelloWorld(ctx context.Context, req *gs.HelloRequest) (rep *gs.HelloResponse, err error) {
	rep = &gs.HelloResponse{
		Resp: req.GetReq(),
	}
	err = nil
	return
}

func main() {
	// 创建服务器监听
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("listen error")
	}
	// 创建grpc实例
	s := grpc.NewServer()
	// 注册方法
	gs.RegisterHelloServer(s, &server{})
	// 开启服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("%v", err)
	}
}
