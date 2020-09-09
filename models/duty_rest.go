package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DutyRest struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	Datetime   string    `xorm:"default '' comment('申请调休日期') VARCHAR(50)"`
	Type       int       `xorm:"default 0 comment('申请调休类型，0：上午，1：下午，2：全天') TINYINT(3)"`
	Proposer   string    `xorm:"default '' comment('申请人') VARCHAR(50)"`
	Checker    string    `xorm:"default '' comment('审核人') VARCHAR(50)"`
	Response   int       `xorm:"comment('审核人的批复，状态 0为默认、1为同意、2为拒绝') TINYINT(1)"`
	CreatedOn  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	ResponseOn time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('审批时间') TIMESTAMP"`
	Backup1    string    `xorm:"default '' VARCHAR(50)"`
	Backup2    string    `xorm:"default '' VARCHAR(50)"`
}

func AddDutyRest(data map[string]interface{}) error {
	rest := DutyRest{
		Datetime:  data["request_time"].(string),
		Type:      data["type"].(int),
		Proposer:  data["proposer"].(string),
		Checker:   data["checker"].(string),
		CreatedOn: time.Now(),
		Response:  0,
	}
	if err := db.Create(&rest).Error; err != nil {
		return err
	}

	return nil
}

func GetAll() ([]DutyRest, error) {
	var (
		rests []DutyRest
		err   error
	)

	err = db.Find(&rests).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return rests, err
}

func CheckIsExist(datetime, proposer string) (bool, error) {
	var (
		rest DutyRest
		err  error
	)
	err = db.Where(DutyRest{
		Datetime: datetime,
		Proposer: proposer,
	}).First(&rest).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if rest.Id > 0 {
		return true, nil
	}

	return false, nil
}

func GetByChecker(checker string) ([]DutyRest, error) {
	var (
		rests []DutyRest
		err   error
	)

	err = db.Where("checker = ? and Response = 0 ", checker).Find(&rests).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return rests, err

}

func EditRest(id int, data interface{}) error {
	if err := db.Model(&DutyRest{}).Where("id = ? ", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAll() error {
	var (
		rests []DutyRest
		err   error
	)

	err = db.Delete(&rests).Error

	if err != nil {
		return err
	}

	return nil

}

func DeleteByName(name string) error {
	var (
		rests []DutyRest
		err   error
	)
	err = db.Where("proposer = ?", name).Delete(&rests).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteById(id int) error {
	var (
		rest DutyRest
		err  error
	)

	err = db.Where("id = ?", id).Delete(&rest).Error

	if err != nil {
		return err
	}

	return nil
}

func GetRestByName(name string, state int) ([]DutyRest, error) {
	var (
		rests []DutyRest
		err   error
	)
	if state == 0 {
		err = db.Where("proposer = ? and response = 0", name).Find(&rests).Error
	} else {
		err = db.Where("proposer = ? and response != 0", name).Find(&rests).Error
	}

	if err != nil {
		return nil, err
	}

	return rests, nil

}
func GetRestById(id int) (DutyRest, error) {
	var (
		rest DutyRest
		err  error
	)

	err = db.Where("id = ?", id).Find(&rest).Error

	if err != nil {
		return DutyRest{}, err
	}

	return rest, nil

}

func GetRestByDay(dateTime string) ([]DutyRest, error) {
	var (
		rests []DutyRest
		err   error
	)

	err = db.Where("dateTime = ? and response = 1", dateTime).Find(&rests).Error

	if err != nil {
		return nil, err
	}

	return rests, nil

}

func AgreeMorningAndFullDay(dateTime string, data interface{}) error {
	if err := db.Model(&DutyRest{}).Where("datetime = ? and type != 1 and response = 0", dateTime).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AgreeAfternoon(dateTime string, data interface{}) error {
	if err := db.Model(&DutyRest{}).Where("datetime = ? and type == 1 and response = 0", dateTime).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
