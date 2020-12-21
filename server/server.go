package server

import (
	"github.com/liujunren93/tcc/server/model"
	"gorm.io/gorm"
	"time"
)

type TransactionManage struct {
	option options
}

type option func(*options)

var transactionManage *TransactionManage

func NewTransaction(opts ...option) (*TransactionManage, error) {
	if transactionManage != nil {
		return transactionManage, nil
	}
	var defaultManage TransactionManage
	for _, opt := range opts {
		opt(&defaultManage.option)
	}
	err := initMigrate(defaultManage.option.DB)
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

//LogTry 记录try
func (t *TransactionManage) LogTry(endpointList []*model.Endpoint) error {

	return t.option.DB.Create(&endpointList).Error
}

//LogCancel  记录Cancel
func (t *TransactionManage) LogCancel(endpointList []*model.Endpoint) error {

	return t.option.DB.Updates(&endpointList).Error
}

//LogConfirm 记录Confirm
func (t *TransactionManage) LogConfirm(endpointList []*model.Endpoint) error {
	return t.option.DB.Updates(&endpointList).Error
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


