package models

import (
	"time"
)

type DutyExchange struct {
	Id            int       `xorm:"not null pk autoincr INT(10)"`
	RequestTime   string    `xorm:"default '' comment('申请日期') VARCHAR(50)"`
	Proposer      string    `xorm:"default '' comment('申请人') VARCHAR(50)"`
	Respondent    string    `xorm:"default '' comment('被申请对象') VARCHAR(50)"`
	RequestedTime string    `xorm:"default '' comment('被申请交换日期') VARCHAR(50)"`
	Response      int       `xorm:"comment('被申请对象的回应，状态 0为拒绝、1为同意') TINYINT(1)"`
	CreatedOn     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	ResponseOn    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('回应时间') TIMESTAMP"`
	Backup1       string    `xorm:"default '' VARCHAR(50)"`
	Backup2       string    `xorm:"default '' VARCHAR(50)"`
}

func AddDutyExchange(data map[string]interface{}) error {
	exchange := DutyExchange{
		RequestTime:   data["request_time"].(string),
		Proposer:      data["proposer"].(string),
		Respondent:    data["respondent"].(string),
		Response:      data["response"].(int),
		RequestedTime: data["requested_time"].(string),
		CreatedOn:     time.Now(),
	}
	if err := db.Create(&exchange).Error; err != nil {
		return err
	}

	return nil
}

func GetExchangeAll() ([]DutyExchange, error) {
	var (
		exchanges []DutyExchange
		err       error
	)
	if err = db.Find(&exchanges).Error; err != nil {
		return nil, err
	}

	return exchanges, nil
}
