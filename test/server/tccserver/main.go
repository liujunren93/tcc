package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
	"github.com/liujunren93/share/server"
	proto1 "github.com/liujunren93/tcc/proto"
	"google.golang.org/grpc"
)

func main()  {
	newRegistry := etcd.NewRegistry()
	 newRegistry.Init(registry.WithAddrs("127.0.0.1:2379"))
	grpcServer := server.NewGrpcServer(server.WithName("tcc.registry"))
	grpcServer.Registry(newRegistry)
	proto1.RegisterTccServerServer(grpcServer.Server().(*grpc.Server),new(server1))
	grpcServer.Run()

}

type  server1 struct {

}

func (s server1) Try(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	panic("implement me")
}

func (s server1) Cancel(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	panic("implement me")
}

func (s server1) Confirm(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	panic("implement me")
}

func (s server1) Registry(ctx context.Context, req *proto1.RegistryReq) (*proto1.RegistryRes, error) {
	var res proto1.RegistryRes
	res.TxID=1
	fmt.Println("register")
	return &res,nil
}

func (s server1) AddEndpoint(ctx context.Context, req *proto1.AddEndpointReq) (*proto1.TccRes, error) {
	panic("implement me")
}







