package main

import (
	"context"
	gs "go_rpc/go_rpc"
	"log"
	"strconv"
	"time"

	grpc "google.golang.org/grpc"
)

func main() {
	// 创建服务器连接
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("dial server error")
	}
	// 连接服务
	client := gs.NewHelloClient(conn)
	// 每隔一秒调用一次
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		// 调用方法
		resp, err := client.HelloWorld(context.Background(), &gs.HelloRequest{Req: "client send:" + strconv.Itoa(t.Second())})
		if err != nil {
			log.Println("%v", err)
		}
		// 输出返回值
		log.Printf("%v", "client receive:"+resp.GetResp())
	}
}
