package server

import (
	"github.com/liujunren93/tcc/model"
	"github.com/liujunren93/tcc/proto"
	"gorm.io/gorm"
	"sync"
	"time"
)

type TransactionManage struct {
	DB *gorm.DB
}

var initMigrateSyncOne sync.Once

var transactionManage *TransactionManage

func NewTransaction(db *gorm.DB) (t *TransactionManage, err error) {
	if transactionManage != nil {
		return transactionManage, nil
	}
	transactionManage=new(TransactionManage)
	transactionManage.DB = db
	initMigrateSyncOne.Do(func() {
		err = initMigrate(transactionManage.DB)

	})
	return transactionManage, err
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
	err := t.DB.Create(&data).Error
	return data.ID, err
}

//Log
// 记录节点事务
func (t *TransactionManage) Log(db **gorm.DB, endpoint []*proto.LogActionData) error {
	tx := *db

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

//ListTransaction
// 事务列表
//0：待处理（默认），1：提交成功，2：提交失败（需要继续提交），3：回滚成功，4：回滚失败（需要继续回滚），5：人工干预
func (t *TransactionManage) ListTransaction(status, page, pageSize int) (list []*model.Transaction, total int64) {
	db := t.DB
	if status != 0 {
		db = db.Where("status=?", status)
	}
	db.Count(&total)
	db.Limit(pageSize).Offset(pageSize * page).Find(&list)
	return
}
