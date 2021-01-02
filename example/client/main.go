package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	sclient "github.com/liujunren93/share/client"
	"github.com/liujunren93/share/registry"
	"github.com/liujunren93/share/registry/etcd"
	"github.com/liujunren93/tcc/client"
	"github.com/liujunren93/tcc/example/db"
	"github.com/liujunren93/tcc/example/proto"
	"google.golang.org/grpc"
)

var map1 = map[string]string{
	"name": "liujunren",
	"age":  "18",
}
var map2 = map[string]string{
	"name": "lili",
	"age":  "18",
	"hh":   "",
}
var thisClient *sclient.Client

func init() {
	initClient()
}

func initClient() {
	thisClient = sclient.NewClient()
	r := etcd.NewRegistry()
	r.Init(registry.WithAddrs("127.0.0.1:2379"))
	thisClient.Init(sclient.WithRegistry(r))
}

func main() {
	buy()
}

func buy() error {
	var err error
	server := getCentServer()
	tccClient, err := client.NewTccClient(client.Config{
		DB:         db.DB(),
		TccCenter: server,
		ClientName: "test",
		Ctx:        nil,
		Logger:     nil,
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			tccClient.Cancel(context.TODO())
		} else {
			tccClient.Confirm(context.TODO())
		}
	}()
	try, err := tccClient.Try(client.NewTccOption(getClientServer1(), "client1", map1), context.TODO())
	if err != nil {
		return err
	}
	map2["hh"] = try.Data["hh"]
	res, err := tccClient.Try(client.NewTccOption(getClientServer2(), "client2", map2), context.TODO())
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
func getCentServer() *grpc.ClientConn {
	next, err := thisClient.Client("tcc.registry")
	if err != nil {
		return nil
	}
	conn, err := next()
	return conn
}

func getClientServer1() proto.Server1TccClient {
	next, err := thisClient.Client("tcc.server1")
	if err != nil {
		return nil
	}
	conn, err := next()
	return proto.NewServer1TccClient(conn)
}
func getClientServer2() proto.Server1TccClient {
	next, err := thisClient.Client("tcc.server2")
	if err != nil {
		return nil
	}
	conn, err := next()
	// 获取grpc client
	return proto.NewServer2TccClient(conn)
}
