package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DutyExchange struct {
	Id            int       `xorm:"not null pk autoincr INT(10)"`
	RequestTime   string    `xorm:"default '' comment('申请日期') VARCHAR(50)"`
	Proposer      string    `xorm:"default '' comment('申请人') VARCHAR(50)"`
	Respondent    string    `xorm:"default '' comment('被申请对象') VARCHAR(50)"`
	RequestedTime string    `xorm:"default '' comment('被申请交换日期') VARCHAR(50)"`
	Response      int       `xorm:"comment('被申请对象的回应，状态 0为未回应、1为同意、2为拒绝') TINYINT(1)"`
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

func AddExchange(requestTime, proposer, Respondent, RequestedTime string, response int) error {
	exchange := DutyExchange{
		RequestTime:   requestTime,
		Proposer:      proposer,
		Respondent:    Respondent,
		RequestedTime: RequestedTime,
		Response:      0,
	}
	if err := db.Create(&exchange).Error; err != nil {
		return err
	}

	return nil
}

func IsExistDay(requestTime, requestedTime string) (bool, error) {
	var exchange DutyExchange
	err := db.Select("id").Where("request_time = ? or requested_time = ? ", requestTime, requestedTime).First(&exchange).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if exchange.Id > 0 {
		return true, nil
	}

	return false, nil

}

func DeleteExchangeAll() error {
	if err := db.Delete(&DutyExchange{}).Error; err != nil {
		return err
	}
	return nil
}

func GetMyExchange(proposer string, response int) ([]DutyExchange, error) {
	var (
		exchanges []DutyExchange
		err       error
	)
	if response == 0 {
		if err = db.Where("proposer = ? and response = 0", proposer).Find(&exchanges).Error; err != nil {
			return nil, err
		}
	} else {
		if err = db.Where("proposer = ? and response != 0", proposer).Find(&exchanges).Error; err != nil {
			return nil, err
		}
	}

	return exchanges, nil
}

func GetMyExamineExchange(respondent string, response int) ([]DutyExchange, error) {
	var (
		exchanges []DutyExchange
		err       error
	)
	if response == 0 {
		if err = db.Where("respondent = ? and response = 0", respondent).Find(&exchanges).Error; err != nil {
			return nil, err
		}
	} else {
		if err = db.Where("respondent = ? and response != 0", respondent).Find(&exchanges).Error; err != nil {
			return nil, err
		}
	}

	return exchanges, nil
}

func GetExchangeById(id int) (DutyExchange, error) {
	var (
		exchange DutyExchange
		err      error
	)

	err = db.Where("id = ?", id).Find(&exchange).Error

	if err != nil {
		return DutyExchange{}, err
	}

	return exchange, nil

}
func DeleteExchangeById(id int) error {
	var (
		exchange DutyExchange
		err      error
	)

	err = db.Where("id = ?", id).Delete(&exchange).Error

	if err != nil {
		return err
	}

	return nil
}

func EditExchange(id int, data interface{}) error {
	if err := db.Model(&DutyExchange{}).Where("id = ? ", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
