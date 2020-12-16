package model

import "gorm.io/gorm"

type Action struct {
	gorm.Model
	serverName string `gorm:"server_name;type:server_name;not null;default:'';comment:'服务名'"`
	TxID       uint   `gorm:"tx_id;type:int;not null;default:0;comment:'主事务id'"`
	ParamData  string `gorm:"param_data;type:varchar(4000);not null;default:'';comment:'一阶段方法参数数据'"`
	Status     int    `gorm:"status;type:int;not null;default:0;comment:'0：待处理（默认），1：提交成功，2：提交失败（需要继续提交），3：回滚成功，4：回滚失败（需要继续回滚），5：人工干预'"`
	RetryCount int    `gorm:"retry_count;type:int;not null;default:0;comment:'当前状态重试次数'"`
}

//Create 创建本地事务
func (a *Action) Create(DB *gorm.DB) error {
	return DB.Create(&a).Error
}

//Create 创建本地事务
func (a *Action) Update(DB *gorm.DB, status, RetryCount int) error {
	return DB.Model(a).Where("id=?", a.ID).Updates(map[string]int{
		"status":      status,
		"retry_count": RetryCount,
	}).Error
}
