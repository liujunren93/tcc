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
	Option options
	conf   Conf
}
type Conf struct {
	Logger *logrus.Logger
}

type option func(options *options)

var syncOnce sync.Once

//NewTccClient 初始化tcc
func NewTccClient( opts ...option) (*client, error) {
	var cli client
	cli.Option = DefaultOptions

	for _, opt := range opts {
		opt(&cli.Option)
	}
	if cli.Option.DB == nil {
		return nil, errors.New("DB must be set")
	}
	initMigrate(cli.Option.DB)
	return &cli, nil
}

//initMigrate 数据迁移
func initMigrate(db *gorm.DB) {
	syncOnce.Do(func() {
		models := []interface{}{
			&model.Action{},
		}
		db.AutoMigrate(models...)
	})
}

//Try
func (c *client) Try(tryList ...tccOption) (err error) {
	var successList []tccOption
	txID, err := c.serverRegistry()
	if err != nil {
		return err
	}
	defer func() {

		if err != nil {
			if successList != nil {
				c.Cancel(successList...)
			}
		} else {
			if successList != nil {
				c.Confirm(successList...)
			}
		}

	}()
	for _, t := range tryList {
		action := model.Action{
			TxID:      uint(txID),
			ParamData: t.paramData.String(),
			Status:    0,
		}
		err = action.Create(c.Option.DB)
		if err != nil {
			return
		}
		t.paramData["tx_id"] = strconv.FormatInt(txID, 10)
		t.paramData["endpoint_tx_id"] = strconv.FormatInt(int64(action.ID), 10)
		t.paramData["param_data"] = t.paramData.String()
		t.paramData["server_name"] = t.serverName
		tryRes, err := t.tcc.Try(t.ctx, &proto.TccReq{Data: t.paramData})
		if err != nil {
			return err
		}
		if tryRes.Code != 200 {
			return errors.New(tryRes.Msg)
		}
		successList = append(successList, t)
	}
	return
}

//Confirm 提交
func (c *client) Confirm(confirmList ...tccOption) (err error) {

	retryCount := 0
	var action model.Action
	for _, t := range confirmList {
	retry://重试
		cancel, err := t.tcc.Confirm(t.ctx, &proto.TccReq{Data: t.paramData})
		if err != nil {
			if retryCount <= 3 {
				retryCount++
				goto retry
			} else {
				dbErr := action.Update(c.Option.DB, 2, 0)
				c.conf.Logger.Errorf("TccConfirmError:data:%v;err:%v \n",action,dbErr)

				return err
			}
		}
		if cancel.Code != 200 {
			if retryCount <= 3 {
				retryCount++
				goto retry
			} else {
				dbErr := action.Update(c.Option.DB, 2, 0)
				c.conf.Logger.Errorf("TccConfirmError:data:%v;err:%v \n",action,dbErr)
				c.conf.Logger.Errorln(dbErr)
				return err
			}
		}
		dbErr := action.Update(c.Option.DB, 1, 0)
		c.conf.Logger.Errorf("TccConfirmError:data:%v;err:%v \n",action,dbErr)
		retryCount = 0
	}
	return nil
}

//Cancel 回滚
func (c *client) Cancel(cancelList ...tccOption) (err error) {

	retryCount := 0
	var action model.Action
	for _, t := range cancelList {
	retry://重试
		cancel, err := t.tcc.Cancel(t.ctx, &proto.TccReq{Data: t.paramData})
		if err != nil {
			if retryCount <= 3 {
				retryCount++
				goto retry
			} else {
				dbErr := action.Update(c.Option.DB, 4, 0)
				c.conf.Logger.Errorf("TccConfirmError:data:%v;err:%v \n",action,dbErr)
				return err
			}
		}
		if cancel.Code != 200 {
			if retryCount <= 3 {
				retryCount++
				goto retry
			} else {
				dbErr := action.Update(c.Option.DB, 4, 0)
				c.conf.Logger.Errorf("TccConfirmError:data:%v;err:%v \n",action,dbErr)
				return err
			}
		}
		dbErr := action.Update(c.Option.DB, 3, 0)
		c.conf.Logger.Errorf("TccConfirmError:data:%v;err:%v \n",action,dbErr)
		retryCount = 0
	}
	return nil
}


//registry 在中心服务器注册获取主事务id
func (c *client) serverRegistry() (txID int64, err error) {

	tccClient := proto.NewTccServerClient(c.Option.GrpcClient)
	ctx := c.Option.Ctx
	if c.Option.Ctx == nil {
		ctx = context.TODO()
	}
	registry, err := tccClient.Registry(ctx, &proto.RegistryReq{})

	if err != nil {
		return 0, err
	}
	return registry.TxID, nil
}

//ServerTry 将信息提交的中心服务器
func (c *client) serverTry(data map[string]string) (err error) {

	tccClient := proto.NewTccServerClient(c.Option.GrpcClient)
	ctx := c.Option.Ctx
	if c.Option.Ctx == nil {
		ctx = context.TODO()
	}
	res, err := tccClient.Try(ctx, &proto.TccReq{Data: data})
	if err != nil {
		return err
	}

	if res.Code != 200 {
		return errors.New(res.Msg)
	}
	return nil
}
