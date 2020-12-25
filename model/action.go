package model

import "gorm.io/gorm"

//Action 客户端记录事务信息
type Action struct {
	gorm.Model
	TxID       uint   `gorm:"tx_id;type:int;not null;default:0;comment:'主事务id'"`
	Endpoint   string `gorm:"endpoint;type:varchar(255);not null;default:'';comment:'节点'"`
	ParamData  string `gorm:"param_data;type:varchar(4000);not null;default:'';comment:'一阶段方法参数数据'"`
	Level      int    `gorm:"level;type:int;not null;default:0;comment:'0:try,1:confirm,2:cancel,3:人工'"`
	Status     int    `gorm:"status;type:int;not null;default:0;comment:'0:待处理，1：成功；2：失败''"`
	RetryCount int    `gorm:"retry_count;type:int;not null;default:0;comment:'当前状态重试次数'"`
}


