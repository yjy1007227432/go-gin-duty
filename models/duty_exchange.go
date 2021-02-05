package models

import (
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type DutyExchange struct {
	Id            int       `xorm:"not null pk autoincr INT(10)"`
	RequestTime   string    `xorm:"default '' comment('申请日期') VARCHAR(50)"`
	Proposer      string    `xorm:"default '' comment('申请人') VARCHAR(50)"`
	Respondent    string    `xorm:"default '' comment('被申请对象') VARCHAR(50)"`
	RequestedTime string    `xorm:"default '' comment('被申请交换日期') VARCHAR(50)"`
	Response      int       `xorm:"comment('被申请对象的回应，状态 0为未回应、1为同意、2为拒绝') TINYINT(1)"`
	ExchangeType  int       `xorm:"comment('换班类型，1，晚班，2,周末白班，3，crm工作日特殊班'，4，周末全天班) TINYINT(1)"`
	CreatedOn     time.Time `xorm:" comment('创建时间') TIMESTAMP"`
	ResponseOn    time.Time `xorm:" comment('回应时间') TIMESTAMP"`
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

func UpdateResponseExchange(id, response int) error {
	var err error

	exchange := DutyExchange{
		Id: id,
	}
	if err = db.Model(&exchange).Update("response", response).Error; err != nil {
		return err
	}
	return nil
}

func GetExchangeByDate(nowDay string) ([]DutyExchange, error) {
	var (
		exchanges []DutyExchange
		err       error
	)
	if err = db.Where("request_time = ? or  requested_time = ? and response = 0", nowDay).Find(&exchanges).Error; err != nil {
		return nil, err
	}

	return exchanges, nil
}

func AddExchange(requestTime, proposer, Respondent, RequestedTime string, ExchangeType int) error {
	exchange := DutyExchange{
		RequestTime:   requestTime,
		Proposer:      proposer,
		Respondent:    Respondent,
		RequestedTime: RequestedTime,
		ExchangeType:  ExchangeType,
		Response:      0,
		CreatedOn:     time.Now(),
		ResponseOn:    time.Now(),
	}
	if err := db.Create(&exchange).Error; err != nil {
		return err
	}

	return nil
}

func IsExistDay(requestTime, requestedTime, respondent, proposer string) (bool, error) {
	var exchange DutyExchange
	err := db.Select("id").Where("((request_time = ? && proposer = ?) or (request_time = ? && proposer = ?) or (requested_time = ? && respondent = ?) or (requested_time = ? && respondent = ?)) && response =0  ", requestedTime, proposer, requestedTime, respondent, requestedTime, proposer, requestedTime, respondent).First(&exchange).Error
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
	if response == -1 {
		if err = db.Where("proposer = ? ", proposer).Find(&exchanges).Error; err != nil {
			return nil, err
		}
	} else {
		if err = db.Where("proposer = ? and response = ?", proposer, response).Find(&exchanges).Error; err != nil {
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
	if response == -1 {
		if err = db.Where("respondent = ? ", respondent).Find(&exchanges).Error; err != nil {
			return nil, err
		}
	} else {
		if err = db.Where("respondent = ? and response = ?", respondent, response).Find(&exchanges).Error; err != nil {
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

func ExchangeTwo(id int, group string, data interface{}) error {
	response := data.(map[string]interface{})["response"].(int)
	exchange := DutyExchange{}
	err := db.Where("id = ?", id).Find(&exchange).Error
	rotaRequest := DutyRota{}
	//查询申请换班人的值班日期值班表
	err = db.Where("datetime = ?", exchange.RequestTime).First(&rotaRequest).Error
	rotaRequested := DutyRota{}
	//查询被申请换班人的值班日期值班表
	err = db.Where("datetime = ?", exchange.RequestedTime).First(&rotaRequested).Error

	//事务操作
	conn := db.Begin()
	//同意换班之后更新值班表
	err = db.Model(&DutyExchange{}).Where("id = ? ", id).Updates(data).Error
	err = db.Where("id = ?", id).First(&exchange).Error

	if err != nil {
		return err
	} else if response == 2 {
		if group == "crm" {
			if exchange.ExchangeType == 1 || exchange.ExchangeType == 4 {
				rotaRequest.CrmLate = strings.Replace(rotaRequest.CrmLate, exchange.Proposer, exchange.Respondent, 1)
				if exchange.ExchangeType == 1 {
					rotaRequested.CrmLate = strings.Replace(rotaRequested.CrmLate, exchange.Respondent, exchange.Proposer, 1)
				}
			}
			if exchange.ExchangeType == 2 || exchange.ExchangeType == 4 {
				rotaRequest.CrmWeekendDay = strings.Replace(rotaRequest.CrmWeekendDay, exchange.Proposer, exchange.Respondent, 1)
				if exchange.ExchangeType == 1 {
					rotaRequested.CrmWeekendDay = strings.Replace(rotaRequested.CrmWeekendDay, exchange.Respondent, exchange.Proposer, 1)
				}
			}
			if exchange.ExchangeType == 3 {
				rotaRequest.CrmDutySpecial = strings.Replace(rotaRequest.CrmDutySpecial, exchange.Proposer, exchange.Respondent, 1)
				if exchange.ExchangeType == 1 {
					rotaRequested.CrmDutySpecial = strings.Replace(rotaRequested.CrmDutySpecial, exchange.Respondent, exchange.Proposer, 1)
				}
			}
		} else {
			rotaRequest.BillingLate = strings.Replace(rotaRequest.BillingLate, exchange.Proposer, exchange.Respondent, 1)
			rotaRequest.BillingWeekendDay = strings.Replace(rotaRequest.BillingWeekendDay, exchange.Proposer, exchange.Respondent, 1)
			if exchange.ExchangeType == 1 {
				rotaRequested.BillingLate = strings.Replace(rotaRequested.BillingLate, exchange.Respondent, exchange.Proposer, 1)
				rotaRequested.BillingWeekendDay = strings.Replace(rotaRequested.BillingWeekendDay, exchange.Respondent, exchange.Proposer, 1)
			}
		}
	}
	err = db.Model(&DutyRota{}).Update(&rotaRequest).Error
	err = db.Model(&DutyRota{}).Update(&rotaRequested).Error

	//更新申请表
	exchange.Response = response
	err = db.Model(&exchange).Update(exchange).Error

	if err != nil {
		return err
	}
	//提交事务
	conn.Commit()
	return nil
}
