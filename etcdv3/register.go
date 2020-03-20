package etcdv3

import (
	"context"
	"log"
	"time"

	clientv3 "github.com/coreos/etcd/clientv3"
)

type Service struct {
	client *clientv3.Client
}

func NewService() (service *Service) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	defer cli.Close()
	if err != nil {
		log.Fatal(err)
	}
	service = &Service{
		client: cli,
	}
	return service
}

// 注册服务
func (s *Service) Register(service, ip string) {
	// 5秒超时
	lease, err := s.client.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	s.client.Put(context.TODO(), service, ip, clientv3.WithLease(lease.ID))
	// 自动续约
	s.client.KeepAlive(context.TODO(), lease.ID)
}

// 获取服务
func (s *Service) GetService(service string) (resp *clientv3.GetResponse) {
	resp, _ = s.client.Get(context.TODO(), service)
	return
}
