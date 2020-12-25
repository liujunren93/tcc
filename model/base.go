package model

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model
	TxID       uint `gorm:"tx_id;type:int;not null;default:0;comment:'主事务id'"`
	EndpointID uint `gorm:"endpoint_id;type:int;not null;default:0;comment:'节点事务id'"`
	TccLevel   int  `gorm:"tcc_level;type:int;not null;default:0;comment:'0:try,1:confirm,2:cancel,3:人工'"`
	TccStatus  int  `gorm:"tcc_status;type:int;not null;default:0;comment:'0:待处理，1：成功；2：失败''"`
}
