package model

import "gorm.io/gorm"

//Transaction 事务中心
type Transaction struct {
	gorm.Model
	BeginTime int64 `gorm:"begin_time;type:int;not null;default:0;comment:'开始时间'"`
	EndTime   int64 `gorm:"end_time;type:int;not null;default:0;comment:'结束时间'"`
	Level     int   `gorm:"level;type:int;not null;default:0;comment:'0:try,1:confirm,2:cancel,3:人工'"`
	Status    int   `gorm:"status;type:int;not null;default:0;comment:'0:待处理，1：成功；2：失败''"`
}

//Endpoint 记录节点事务信息
type Endpoint struct {
	gorm.Model
	TxID         uint   `gorm:"tx_id;type:int;not null;uniqueIndex:tx_endpoint;default:0;comment:'主事务id'"`
	EndpointTxID uint   `gorm:"endpoint_tx_id;type:int;not null;uniqueIndex:tx_endpoint;default:0;;comment:'节点事务id'"`
	Endpoint     string `gorm:"endpoint;type:varchar(255);not null;default:'';comment:'节点'"`
	Level        int    `gorm:"level;type:int;not null;default:0;comment:'0:try,1:confirm,2:cancel,3:人工'"`
	Status       int    `gorm:"status;type:int;not null;default:0;comment:'0:待处理，1：成功；2：失败''"`
	RetryCount   int    `gorm:"retry_count;type:int;not null;default:0;comment:'当前状态重试次数'"`
}
