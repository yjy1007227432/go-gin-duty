package models

import (
	"time"
)

type DutyRota struct {
	Id                 int       `xorm:"not null pk autoincr INT(10)"`
	Datetime           string    `xorm:"default '' comment('日期') VARCHAR(50)"`
	Week               string    `xorm:"default '' comment('星期') VARCHAR(50)"`
	BillingLate        string    `xorm:"default '' comment('计费晚班人员') VARCHAR(50)"`
	BillingWeekendLate string    `xorm:"default '' comment('计费周末晚班人员') VARCHAR(50)"`
	CrmLate            string    `xorm:"default '' comment('crm晚班人员') VARCHAR(50)"`
	CrmWeekendLate     string    `xorm:"default '' comment('crm周末晚班人员') VARCHAR(50)"`
	CrmDuty            string    `xorm:"default '' comment('crm值班人员') VARCHAR(50)"`
	CreatedOn          time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	CreatedBy          string    `xorm:"default '' comment('创建人') VARCHAR(100)"`
	ModifiedOn         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	ModifiedBy         string    `xorm:"default '' comment('修改人') VARCHAR(255)"`
	Backup1            string    `xorm:"default '' VARCHAR(50)"`
	Backup2            string    `xorm:"default '' VARCHAR(50)"`
}
