package server

import (
	"errors"
	"github.com/liujunren93/tcc/proto"
	"github.com/liujunren93/tcc/server/model"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

type TransactionManage struct {
	option options
}

var initMigrateSyncOne sync.Once

type option func(*options)

var transactionManage *TransactionManage

func NewTransaction(opts ...option) (t *TransactionManage, err error) {
	if transactionManage != nil {
		return transactionManage, nil
	}
	var defaultManage TransactionManage
	for _, opt := range opts {
		opt(&defaultManage.option)
	}
	initMigrateSyncOne.Do(func() {
		err = initMigrate(defaultManage.option.DB)

	})
	return &defaultManage, err
}

func initMigrate(db *gorm.DB) error {

	models := []interface{}{
		&model.Endpoint{},
		&model.Transaction{},
	}
	return db.AutoMigrate(models...)
}

//Registry 注册事务获取事务id
func (t *TransactionManage) Registry() (uint, error) {
	var data model.Transaction
	data.BeginTime = time.Now().Local().Unix()
	err := t.option.DB.Create(&data).Error
	return data.ID, err
}

//Log
// 记录节点事务
func (t *TransactionManage) Log(endpoint []*proto.LogActionData) error {
	tx := t.option.DB
	endpointList, transaction := logActionData2Model(endpoint)
	for _, m := range endpointList {
		var tmp model.Endpoint
		 tx.Where("tx_id=? and endpoint_tx_id=? and level<?", m.TxID, m.EndpointTxID, m.Level).First(&tmp)
		if tmp.ID != 0 {
			err := tx.Where("id=?", tmp.ID).Model(&model.Endpoint{}).Updates(m).Error
			if err != nil {
				return err
			}
		} else {
			err := tx.Model(&model.Endpoint{}).Create(m).Error
			if err != nil {
				return err
			}
		}

	}
	err := tx.Where("id=?", transaction.ID).Updates(&transaction).Error
	return err

}

//LogActionData2Model
func logActionData2Model(logList []*proto.LogActionData) (endpointList []*model.Endpoint, transaction *model.Transaction) {
	var t model.Transaction
	for _, log := range logList {
		endpoint := model.Endpoint{
			TxID:         uint(log.TxID),
			EndpointTxID: uint(log.EndpointTxID),
			Endpoint:     log.EndpointName,
			Level:        int(log.Level),
			Status:       int(log.Status),
		}
		t.Level = endpoint.Level
		if endpoint.Status >= t.Status {
			t.Status = endpoint.Status
		}
		t.ID = endpoint.TxID
		endpointList = append(endpointList, &endpoint)
	}
	return endpointList, &t
}

//map2Endpoint
//level 全局执行进度
//status 全局执行状态
func map2Endpoint1(list []map[string]string) (endpointList []*model.Endpoint, transaction *model.Transaction, err error) {
	var res []*model.Endpoint
	var t model.Transaction
	for _, m := range list {
		var endpoint model.Endpoint
		if val, ok := m["tx_id"]; ok {
			txID, _ := strconv.Atoi(val)
			endpoint.TxID = uint(txID)
		} else {
			return nil, nil, errors.New("tx_id can not be empty")
		}
		if val, ok := m["endpoint_tx_id"]; ok {
			endpointTxId, _ := strconv.Atoi(val)
			endpoint.EndpointTxID = uint(endpointTxId)
		} else {
			return nil, nil, errors.New("endpoint_tx_id can not be empty")
		}
		if val, ok := m["status"]; ok {
			status, _ := strconv.Atoi(val)
			endpoint.Status = status
		}
		if val, ok := m["level"]; ok {
			level, _ := strconv.Atoi(val)
			endpoint.Level = level
		}
		if val, ok := m["pk"]; ok {
			pk, _ := strconv.Atoi(val)
			endpoint.ID = uint(pk)
		}
		if val, ok := m["param_data"]; ok {
			endpoint.ParamData = val
		}
		t.Level = endpoint.Level
		if endpoint.Status >= t.Status {
			t.Status = endpoint.Status
		}
		t.ID = endpoint.TxID
		res = append(res, &endpoint)
	}

	return res, &t, nil
}

//ListTransaction
// 事务列表
//0：待处理（默认），1：提交成功，2：提交失败（需要继续提交），3：回滚成功，4：回滚失败（需要继续回滚），5：人工干预
func (t *TransactionManage) ListTransaction(status, page, pageSize int) (list []*model.Transaction, total int64) {
	db := t.option.DB
	if status != 0 {
		db = db.Where("status=?", status)
	}
	db.Count(&total)
	db.Limit(pageSize).Offset(pageSize * page).Find(&list)
	return
}
