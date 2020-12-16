package client

import (
	"context"
	"encoding/json"
	"github.com/liujunren93/tcc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type options struct {
	DB         *gorm.DB
	GrpcClient *grpc.ClientConn
	TccServer  *grpc.ClientConn
	ServerName string
	Ctx        context.Context
	Logger     *logrus.Logger
}

var DefaultOptions =options{
	Ctx:        context.TODO(),
	Logger:     logrus.New(),
}
//WithDB 设置数据库
func WithDB(db *gorm.DB) option {

	return func(opt *options) {
		opt.DB = db
	}
}

func WithLogger(log *logrus.Logger) option {
	return func(opt *options) {
		opt.Logger = log
	}
}

//WithServerName 服务名
func WithServerName(ServerName string) option {
	return func(opt *options) {
		opt.ServerName = ServerName
	}
}

//WithGrpcClient  设置grpc client
func WithGrpcClient(client *grpc.ClientConn) option {
	return func(opt *options) {
		opt.GrpcClient = client
	}
}

//paramData 数据
type paramData map[string]string

func (p paramData) String() string {
	marshal, _ := json.Marshal(&p)
	return string(marshal)
}

type tccOption struct {
	serverName string
	tcc        tcc.Tcc
	paramData  paramData
	ctx        context.Context
}

func NewTccOption(tcc tcc.Tcc, serverName string, ParamData map[string]string, ctx context.Context) tccOption {
	return tccOption{
		serverName: serverName,
		tcc:        tcc,
		paramData:  ParamData,
		ctx:        ctx,
	}
}
