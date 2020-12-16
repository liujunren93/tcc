package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
	"github.com/liujunren93/share/server"
	proto1 "github.com/liujunren93/tcc/proto"
	"github.com/liujunren93/tcc/test/proto"

	"google.golang.org/grpc"
)

func main()  {
	newRegistry := etcd.NewRegistry()
	newRegistry.Init(registry.WithAddrs("127.0.0.1:2379"))
	grpcServer := server.NewGrpcServer(server.WithName("tcc.server1"))
	grpcServer.Registry(newRegistry)
	proto.RegisterServer1TccServer(grpcServer.Server().(*grpc.Server),new(server1))
	grpcServer.Run()

}

type  server1 struct {

}

func (s server1) Try(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	fmt.Println("ser1")
	return &proto1.TccRes{Code: 300,Msg: "server1"}, nil
}

func (s server1) Cancel(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	fmt.Println("Cancel1")

	return &proto1.TccRes{Code: 200,Msg: "server1"}, nil
}

func (s server1) Confirm(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	fmt.Println("Confirm1")
	return nil, nil
}









