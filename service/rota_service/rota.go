package rota_service

import (
	"go-gin-duty-master/models"
	"time"
)

type Rota struct {
	Id                 int
	Datetime           string
	Week               string
	BillingLate        string
	BillingWeekendLate string
	CrmLate            string
	CrmWeekendLate     string
	CrmDuty            string
	CreatedOn          time.Time
	CreatedBy          string
	ModifiedOn         time.Time
	ModifiedBy         string
}

func (t *Rota) ExistByDatetime() (bool, error) {
	return models.ExistRotaByDatetime(t.Datetime)
}

func (t *Rota) Add() error {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["datetime"] = t.Datetime
	m["week"] = t.Week
	m["billing_late"] = t.BillingLate
	m["billing_weekend_late"] = t.BillingWeekendLate
	m["crm_late"] = t.CrmLate
	m["crm_weekend_late"] = t.CrmWeekendLate
	m["crm_duty"] = t.CrmDuty
	m["created_by"] = t.CreatedBy
	err := models.AddDutyRest(m)
	return err
}
