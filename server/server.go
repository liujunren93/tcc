package server

import (
	"github.com/liujunren93/tcc/server/model"
	"gorm.io/gorm"
)

type transactionManage struct {
	option options
}

type option func(*options)

func NewTransaction(opts ...option) *transactionManage {
	var defaultManage transactionManage
	for _, opt := range opts {
		opt(&defaultManage.option)
	}
	initMigrate(defaultManage.option.DB)
	return &defaultManage
}

func initMigrate(db *gorm.DB) {
	models := []interface{}{
		&model.Endpoint{},
		&model.Transaction{},
	}
	db.AutoMigrate(models)
}

//Registry 注册事务获取事务id
func (t *transactionManage) Registry() (uint, error) {
	var data model.Transaction
	err := t.option.DB.Create(&data).Error
	return data.ID, err
}

//Begin 开始事务
func (t *transactionManage) Begin(endpointList []*model.Endpoint) error {
	return t.option.DB.Create(&endpointList).Error
}
