package client

import (
	"context"
	"errors"
	"github.com/liujunren93/tcc/client/model"
	"github.com/liujunren93/tcc/proto"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type client struct {
	option  Config
	conf    Conf
	txID    uint
	tccList []*tccOption
}
type Conf struct {
	Logger *logrus.Logger
}

var syncOnce sync.Once

//NewTccClient 初始化tcc
//func NewTccClient(Config) (*client, error) {
func NewTccClient(conf Config) (*client, error) {
	var cli client
	if conf.Ctx == nil {
		conf.Ctx = context.TODO()
	}
	if conf.Logger == nil {
		conf.Logger = logrus.New()
	}
	cli.option = conf
	if cli.option.DB == nil {
		return nil, errors.New("DB must be set")
	}
	syncOnce.Do(func() {
		initMigrate(cli.option.DB)
	})
	err := serverRegistry(&cli)
	if err != nil {
		cli.option.Logger.Error(err)
		return nil, err
	}
	return &cli, nil
}

//initMigrate 初始化数据
func initMigrate(db *gorm.DB) {
	models := []interface{}{
		&model.Action{},
	}
	err := db.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}

//Try
func (c *client) Try(endpoint tccOption, ctx context.Context) (res *proto.TccRes, err error) {
	action := model.Action{
		TxID:      c.txID,
		Endpoint:  endpoint.endpoint,
		ParamData: endpoint.paramData.String(),
	}
	err = c.option.DB.Create(&action).Error
	if err != nil {
		return
	}
	endpoint.paramData["tx_id"] = strconv.FormatInt(int64(c.txID), 10)
	endpoint.paramData["endpoint_tx_id"] = strconv.FormatInt(int64(action.ID), 10)
	endpoint.paramData["param_data"] = endpoint.paramData.String()
	endpoint.paramData["endpoint"] = endpoint.endpoint
	res, err = endpoint.tcc.Try(ctx, &proto.TccReq{Data: endpoint.paramData})
	var status int64
	defer func() {
		err2 := c.tccServerLog(&proto.LogActionData{
			TxID:         int64(c.txID),
			EndpointTxID: int64(action.ID),
			Level:        1,
			Status:       status,
		})
		if err != nil {
			c.option.Logger.Error(err2)
		}
	}()
	if err != nil {
		return res, err
	}

	c.tccList = append(c.tccList, &endpoint)
	return res, nil
}

//Confirm 提交
func (c *client) Confirm(ctx context.Context) (err error) {
	retryCount := 0
	status := 0
	var log []*proto.LogActionData
	for _, t := range c.tccList {
		var paramData = map[string]string{
			"tx_id":          t.paramData["tx_id"],
			"endpoint_tx_id": t.paramData["endpoint_tx_id"],
		}
	retry:
		res, err := t.tcc.Confirm(ctx, &proto.TccReq{Data: paramData})
		if err != nil || res.Code != 200 {
			if retryCount < 3 {
				retryCount++
				status = 2
				goto retry
			} else {
				err := c.Cancel(ctx)
				if err != nil {
					c.option.Logger.Error(err)
					return err
				}
			}
		} else {
			status = 1
		}
		c.option.DB.Where("id=? and level<=1", t.paramData["endpoint_tx_id"]).Updates(&model.Action{
			Level:      1,
			Status:     status,
			RetryCount: retryCount,
		})
		endpointTxId, _ := strconv.Atoi(t.paramData["endpoint_tx_id"])
		log = append(log, &proto.LogActionData{

			TxID:         int64(c.txID),
			EndpointTxID: int64(endpointTxId),
			Level:        1,
			Status:       int64(status),
		})
	}
	return c.tccServerLog(log...)
}

//Cancel 回滚
func (c *client) Cancel(ctx context.Context) error {
	retryCount := 0
	status := 0
	var log []*proto.LogActionData
	for _, t := range c.tccList {
		var paramData = map[string]string{
			"tx_id":          t.paramData["tx_id"],
			"endpoint_tx_id": t.paramData["endpoint_tx_id"],
		}
	retry:
		res, err := t.tcc.Cancel(ctx, &proto.TccReq{
			Data: paramData,
		})
		if (err != nil || res.Code != 200) && retryCount < 3 {
			retryCount++
			status = 2
			goto retry
		} else {
			status = 1
		}
		c.option.DB.Where("id=?", t.paramData["endpoint_tx_id"]).Updates(&model.Action{
			Level:      2,
			Status:     status,
			RetryCount: retryCount,
		})
		endpointTxId, _ := strconv.Atoi(t.paramData["endpoint_tx_id"])
		log = append(log, &proto.LogActionData{
			TxID:         int64(c.txID),
			EndpointTxID: int64(endpointTxId),
			Level:        2,
			Status:       int64(status),
		})

	}
	return c.tccServerLog(log...)
}

//serverRegistry 在事务中心注册获取主事务id
func serverRegistry(c *client) error {
	tccClient := proto.NewTccServerClient(c.option.TccCenter)
	ctx := c.option.Ctx
	if c.option.Ctx == nil {
		ctx = context.TODO()
	}
	res, err := tccClient.Registry(ctx, &proto.Empty{})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Msg)
	}
	c.txID = uint(res.Data.TxID)
	return nil
}

//tccServerLog 将信息提交的中心服务器
func (c *client) tccServerLog(data ...*proto.LogActionData) error {
	tccClient := proto.NewTccServerClient(c.option.TccCenter)
	ctx := c.option.Ctx
	if c.option.Ctx == nil {
		ctx = context.TODO()
	}
	res, err := tccClient.LogAction(ctx, &proto.LogActionReq{Data: data})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Msg)
	}

	return nil
}
