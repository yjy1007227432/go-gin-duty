package models

import "time"

type DutyOverTime struct {
	Id         int       `xorm:"not null pk autoincr INT(10)"`
	Quantity   float64   `xorm:"default 0 comment('加班天数') INT(10)"`
	Proposer   string    `xorm:"default '' comment('申请人') VARCHAR(50)"`
	Reason     string    `xorm:"default '' comment('理由') VARCHAR(1000)"`
	Checker    string    `xorm:"default '' comment('审核人') VARCHAR(50)"`
	Response   int       `xorm:"comment('状态 0为未回应、1为同意、2为拒绝') TINYINT(1)"`
	CreatedOn  time.Time `xorm:" comment('创建时间') TIMESTAMP"`
	ResponseOn time.Time `xorm:" comment('回应时间') TIMESTAMP"`
	Backup1    string    `xorm:"default '' VARCHAR(50)"`
	Backup2    string    `xorm:"default '' VARCHAR(50)"`
}

func AddDutyOverTime(data map[string]interface{}) error {
	dutyOverTime := DutyOverTime{
		Quantity:  data["quantity"].(float64),
		Proposer:  data["proposer"].(string),
		Reason:    data["reason"].(string),
		Response:  0,
		CreatedOn: time.Now(),
	}
	if err := db.Create(&dutyOverTime).Error; err != nil {
		return err
	}

	return nil
}

func UpdateResponseOverTime(id, response int) error {
	var err error

	dutyOverTime := DutyOverTime{
		Id: id,
	}
	if err = db.Model(&dutyOverTime).Update("response", response).Error; err != nil {
		return err
	}
	return nil
}

func GetOverTimeById(id int) (DutyOverTime, error) {
	var (
		dutyOverTime DutyOverTime
		err          error
	)

	err = db.Where("id = ?", id).Find(&dutyOverTime).Error

	if err != nil {
		return DutyOverTime{}, err
	}
	return dutyOverTime, nil
}

func GetMyDutyOverTime(proposer string) ([]DutyOverTime, error) {
	var (
		dutyOverTimes []DutyOverTime
		err           error
	)
	if err = db.Where("proposer = ? ", proposer).Find(&dutyOverTimes).Error; err != nil {
		return nil, err
	}

	return dutyOverTimes, nil
}

func GetOverTimeAll() ([]DutyOverTime, error) {
	var (
		dutyOverTimes []DutyOverTime
		err           error
	)

	err = db.Find(&dutyOverTimes).Error

	if err != nil {
		return nil, err
	}

	return dutyOverTimes, nil

}

func EditOverTime(id int, data interface{}) error {
	if err := db.Model(&DutyOverTime{}).Where("id = ? ", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
