package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	BeginTime int64 `gorm:"begin_time;type:int;not null;default:0;comment:'开始时间'"`
	EndTime   int64 `gorm:"end_time;type:int;not null;default:0;comment:'结束时间'"`
	Status    int   `gorm:"status;type:int;not null;default:0;comment:'0：待处理（默认），1：提交成功，2：提交失败（需要继续提交），3：回滚成功，4：回滚失败（需要继续回滚），5：人工干预'"`
}

type Endpoint struct {
	gorm.Model
	TxID               uint   `gorm:"tx_id;type:int;not null;default:0;comment:'主事务id'"`
	EndpointServerName string `gorm:"endpoint_server_name;type:varchar(255);not null;default:'';comment:'服务名'"`
	EndpointTxID       uint   `gorm:"endpoint_tx_id;type:int;not null;default:0;comment:'分支事务ID'"`
	ParamData          string `gorm:"param_data;type:varchar(4000);not null;default:'';comment:'一阶段方法参数数据'"`
}
