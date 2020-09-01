package models

import (
	"time"
)

type DutyRest struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	Datetime   string    `xorm:"default '' comment('申请调休日期') VARCHAR(50)"`
	Type       int       `xorm:"default 0 comment('申请调休类型，0：上午，1：下午，2：全天') TINYINT(3)"`
	Proposer   string    `xorm:"default '' comment('申请人') VARCHAR(50)"`
	Checker    string    `xorm:"default '' comment('审核人') VARCHAR(50)"`
	Response   int       `xorm:"comment('审核人的批复，状态 0为拒绝、1为同意') TINYINT(1)"`
	CreatedOn  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	ResponseOn time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('审批时间') TIMESTAMP"`
	Backup1    string    `xorm:"default '' VARCHAR(50)"`
	Backup2    string    `xorm:"default '' VARCHAR(50)"`
}
