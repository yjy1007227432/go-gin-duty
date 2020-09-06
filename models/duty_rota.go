package models

import (
	"github.com/jinzhu/gorm"
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

func AddDutyRota(data map[string]interface{}) error {
	rota := DutyRota{
		Datetime:           data["datetime"].(string),
		Week:               data["week"].(string),
		BillingLate:        data["billing_late"].(string),
		BillingWeekendLate: data["billing_weekend_late"].(string),
		CrmLate:            data["crm_late"].(string),
		CreatedOn:          time.Now(),
		CrmWeekendLate:     data["crm_weekend_late"].(string),
		CrmDuty:            data["crm_duty"].(string),
		CreatedBy:          data["created_by"].(string),
	}
	if err := db.Create(&rota).Error; err != nil {
		return err
	}

	return nil
}

func GetMonth(month string) ([]DutyRota, error) {

	var (
		rotas []DutyRota
		err   error
	)

	err = db.Where("datetime like ?", month+"%").Find(&rotas).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return rotas, err

}

func GetRotaByDay(day string) (DutyRota, error) {

	var (
		rota DutyRota
		err  error
	)
	err = db.Where("datetime = ?", day).Find(&rota).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return DutyRota{}, err
	}
	return rota, err
}

func DeleteMonth(month string) error {

	var (
		rotas []DutyRota
		err   error
	)

	err = db.Where("datetime like ?", month+"%").Delete(&rotas).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil

}

func DeleteDay(day string) error {

	var (
		rotas []DutyRota
		err   error
	)

	err = db.Where("datetime = ?", day).Delete(&rotas).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil

}

func ExistRotaByDatetime(datetime string) (bool, error) {
	var rota DutyRota
	err := db.Select("id").Where("datetime = ? ", datetime).First(&rota).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if rota.Id > 0 {
		return true, nil
	}

	return false, nil
}
