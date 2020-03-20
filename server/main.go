package main

import (
	"context"
	"fmt"
	gs "go_rpc/go_rpc"
	"log"
	"net"

	etcd "go_rpc/etcdv3"

	grpc "google.golang.org/grpc"
)

type server struct{}

const (
	ServiceName = "dt_go_rpc"
	Ip          = "127.0.0.1:50051"
)

// 实现HelloWorld方法
func (s *server) HelloWorld(ctx context.Context, req *gs.HelloRequest) (rep *gs.HelloResponse, err error) {
	rep = &gs.HelloResponse{
		Resp: "server send:" + req.GetReq(),
	}
	err = nil
	// 测试客户端输入
	log.Printf("server get:%v", req.GetReq())
	return
}

func main() {
	etcdService := etcd.NewService()
	etcdService.Register(ServiceName, Ip)
	fmt.Println(etcdService.GetService(ServiceName))

	// 创建服务器监听
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("listen error")
	}
	// 创建grpc实例
	s := grpc.NewServer()
	// 注册服务
	gs.RegisterHelloServer(s, &server{})
	// 开启服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("%v", err)
	}
}
