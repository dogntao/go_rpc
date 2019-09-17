package main

import (
	"context"
	gs "go_rpc/go_rpc"
	"log"
	"os"
	"time"

	grpc "google.golang.org/grpc"
)

func main() {
	// 创建和服务器的连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	// 创建客户端存款
	client := gs.NewGreeterClient(conn)
	// 获取用户输入
	name := "dt"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, calcel := context.WithTimeout(context.Background(), time.Second)
	defer calcel()
	// 调用服务
	r, err := client.HelloWorld(ctx, &gs.HelloRequest{Req: name})
	if err != nil {
		log.Fatalf("can not call %v", err)
	}
	log.Printf("greeting:%v", r.GetResp())
}
