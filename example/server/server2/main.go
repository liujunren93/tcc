package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
	"github.com/liujunren93/share/server"
	proto1 "github.com/liujunren93/tcc/proto"

	"github.com/liujunren93/tcc/example/proto"
	"google.golang.org/grpc"
)

func main()  {
	newRegistry := etcd.NewRegistry()
	 newRegistry.Init(registry.WithAddrs("127.0.0.1:2379"))
	grpcServer := server.NewGrpcServer(server.WithName("tcc.server2"))
	grpcServer.Registry(newRegistry)
	proto.RegisterServer2TccServer(grpcServer.Server().(*grpc.Server),new(server1))
	grpcServer.Run()

}

type  server1 struct {

}

func (s server1) Try(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	fmt.Println("server2",req.Data)
	return &proto1.TccRes{Code: 200}, nil
}

func (s server1) Cancel(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	fmt.Println("Cancel2")
	return &proto1.TccRes{Code: 200,Msg: "server2"}, nil
}

func (s server1) Confirm(ctx context.Context, req *proto1.TccReq) (*proto1.TccRes, error) {
	return &proto1.TccRes{Code: 200, Msg: "server1"}, nil
}







