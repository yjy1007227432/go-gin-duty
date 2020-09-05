package models

import (
	"github.com/jinzhu/gorm"
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

func GetVacationAll() ([]DutyVacation, error) {
	var (
		vacations []DutyVacation
		err       error
	)
	err = db.Find(&vacations).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return vacations, err
}

func GetVacationByName(name string) (DutyVacation, error) {
	var (
		vacation DutyVacation
		err      error
	)

	err = db.Where("name = ?", name).Find(&vacation).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return DutyVacation{}, err
	}
	return vacation, err
}

func DeleteVacationAll() error {

	err := db.Delete(&DutyVacation{}).Error

	if err != nil {
		return err
	}
	return nil
}

func DeleteVacationByName(name string) error {

	err := db.Where("name = ? ", name).Delete(&DutyVacation{}).Error

	if err != nil {
		return err
	}
	return nil
}

func EditVacation(id int, data interface{}) error {
	if err := db.Model(&DutyRest{}).Where("id = ? ", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
