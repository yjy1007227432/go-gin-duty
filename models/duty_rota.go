package models

import (
	"github.com/jinzhu/gorm"
	"go-gin-duty-master/service/rota_service"
	"strings"
	"time"
)

type DutyRota struct {
	Id                int       `xorm:"not null pk autoincr INT(10)"`
	Datetime          string    `xorm:"default '' comment('日期') VARCHAR(50)"`
	Week              string    `xorm:"default '' comment('星期') VARCHAR(50)"`
	BillingLate       string    `xorm:"default '' comment('计费晚班人员') VARCHAR(50)"`
	BillingWeekendDay string    `xorm:"default '' comment('计费周末白班人员') VARCHAR(50)"`
	CrmLate           string    `xorm:"default '' comment('crm晚班人员') VARCHAR(50)"`
	CrmWeekendDay     string    `xorm:"default '' comment('crm周末白班人员') VARCHAR(50)"`
	CrmDutySpecial    string    `xorm:"default '' comment('crm工作日特殊班值班人员') VARCHAR(50)"`
	CreatedOn         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	CreatedBy         string    `xorm:"default '' comment('创建人') VARCHAR(100)"`
	ModifiedOn        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	ModifiedBy        string    `xorm:"default '' comment('修改人') VARCHAR(255)"`
	Backup1           string    `xorm:"default '' VARCHAR(50)"`
	Backup2           string    `xorm:"default '' VARCHAR(50)"`
}

func AddDutyRota(data map[string]interface{}) error {
	rota := DutyRota{
		Datetime:          data["datetime"].(string),
		Week:              data["week"].(string),
		BillingLate:       data["billing_late"].(string),
		BillingWeekendDay: data["billing_weekend_late"].(string),
		CrmLate:           data["crm_late"].(string),
		CreatedOn:         time.Now(),
		CrmWeekendDay:     data["crm_weekend_late"].(string),
		CrmDutySpecial:    data["crm_duty"].(string),
		CreatedBy:         data["created_by"].(string),
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

func UpdateRotaByDateTime(dateTime string, data interface{}) error {
	if err := db.Model(&DutyRest{}).Where("datetime = ? ", dateTime).Updates(data).Error; err != nil {
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
