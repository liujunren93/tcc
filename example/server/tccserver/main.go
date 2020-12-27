package main

import (
	"context"
	"fmt"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/tcc/example/db"
	proto1 "github.com/liujunren93/tcc/proto"
	server2 "github.com/liujunren93/tcc/server"
	"google.golang.org/grpc"
)

var t *server2.TransactionManage

func main() {
	t, _ = server2.NewTransaction(db.DB())
	newRegistry := etcd.NewRegistry()
	newRegistry.Init(registry.WithAddrs("127.0.0.1:2379"))
	grpcServer := server.NewGrpcServer(server.WithName("tcc.registry"))
	grpcServer.Registry(newRegistry)
	proto1.RegisterTccServerServer(grpcServer.Server().(*grpc.Server), new(server1))
	grpcServer.Run()

}

type server1 struct {
}

func (s server1) LogAction(ctx context.Context, req *proto1.LogActionReq) (*proto1.TccRes, error) {
	err := t.Log(req.Data)
	if err != nil {
		fmt.Println(err)
	}
	return &proto1.TccRes{Code: 200}, err
}

func (s server1) Registry(ctx context.Context, req *proto1.Empty) (*proto1.RegistryRes, error) {
	var res proto1.RegistryRes
	res.Code = 200
	u, _ := t.Registry()
	res.Data = &proto1.RegistryResData{TxID: int64(u)}
	return &res, nil
}
