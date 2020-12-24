package client

import (
	"context"
	"encoding/json"
	"github.com/liujunren93/tcc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	DB         *gorm.DB
	TccCenter  *grpc.ClientConn
	ClientName string
	Ctx        context.Context
	Logger     *logrus.Logger
}

//paramData 数据
type paramData map[string]string

func (p paramData) String() string {
	marshal, _ := json.Marshal(&p)
	return string(marshal)
}

//tccOption 发起tcc请求数据
type tccOption struct {
	endpoint  string
	tcc       tcc.Tcc
	paramData paramData
	//ctx       context.Context
}

//NewTccOption
func NewTccOption(tcc tcc.Tcc, endpoint string, paramData map[string]string) tccOption {
	return tccOption{
		endpoint:  endpoint,
		tcc:       tcc,
		paramData: paramData,
		//ctx:       ctx,
	}
}
