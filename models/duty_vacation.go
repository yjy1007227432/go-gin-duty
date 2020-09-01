package models

import (
	"time"
)

type DutyVacation struct {
	Id                   int       `xorm:"not null pk autoincr INT(10)"`
	Name                 string    `xorm:"default '' comment('姓名') VARCHAR(50)"`
	RemainVacation       int       `xorm:"default 0 comment('剩余调休天数') INT(10)"`
	RemainAnnualVacation int       `xorm:"default 0 comment('剩余年休天数') INT(10)"`
	UpdateTime           time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	Backup1              string    `xorm:"default '' VARCHAR(50)"`
	Backup2              string    `xorm:"default '' VARCHAR(50)"`
}
