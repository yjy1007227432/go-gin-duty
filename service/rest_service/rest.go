package rest_service

import (
	"go-gin-duty-master/models"
	"time"
)

type Rest struct {
	Id           int       `form:"id"  json:"id"`
	Datetime     string    `form:"datetime"  json:"datetime" `
	Type         int       ` form:"type"   json:"type"  `
	VacationType int       ` form:"vacation_type"   json:"vacation_type"  `
	Proposer     string    `form:"proposer"   json:"proposer" `
	Checker      string    `form:"checker"  json:"checker" `
	Response     int       `form:"response"  json:"response" `
	CreatedOn    time.Time `form:"created_on"  json:"created_on"  `
	ResponseOn   time.Time `form:"response_on"  json:"response_on" `
	State        int       `form:"state"   json:"state"`
}

func (t *Rest) Add() error {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["type"] = t.Type
	m["request_time"] = t.Datetime
	m["proposer"] = t.Proposer
	m["checker"] = t.Checker
	m["vacation_type"] = t.VacationType
	err := models.AddDutyRest(m)
	return err
}

func (t *Rest) CheckIsExist() (bool, error) {
	IsExist, err := models.CheckIsExist(t.Datetime, t.Proposer, t.Type)
	return IsExist, err
}

func (t *Rest) GetAll() ([]models.DutyRest, error) {
	rests, err := models.GetAll()
	return rests, err
}

func (t *Rest) GetAllowedRestByMonth() ([]models.DutyRest, error) {
	rests, err := models.GetAllowedRestByMonth(t.Datetime)
	return rests, err
}

func (t *Rest) GetRestById() (models.DutyRest, error) {
	rest, err := models.GetRestById(t.Id)
	return rest, err
}

func (t *Rest) GetRestByDayAgree() ([]models.DutyRest, error) {
	rests, err := models.GetRestByDayAgree(t.Datetime)
	return rests, err
}

func (t *Rest) AgreeMorningAndFullDay() error {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["response"] = 2
	err := models.AgreeMorningAndFullDay(t.Datetime, m)
	return err
}

func (t *Rest) AgreeAfternoon() error {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["response"] = 2
	err := models.AgreeAfternoon(t.Datetime, m)
	return err
}

func (t *Rest) GetByChecker() ([]models.DutyRest, error) {
	rests, err := models.GetByChecker(t.Checker)
	return rests, err
}

func (t *Rest) Edit() error {
	data := make(map[string]interface{})
	data["response"] = t.Response

	return models.EditRest(t.Id, data)
}

func (t *Rest) DeleteAll() error {
	err := models.DeleteAll()
	return err
}

func (t *Rest) GetRestsByNameState() ([]models.DutyRest, error) {
	rests, err := models.GetRestByNameState(t.Proposer, t.State)
	return rests, err
}
func (t *Rest) GetRestsByName() ([]models.DutyRest, error) {
	rests, err := models.GetRestByName(t.Proposer)
	return rests, err
}

func (t *Rest) DeleteById() error {
	err := models.DeleteById(t.Id)
	return err
}

func (t *Rest) DeleteByName() error {
	err := models.DeleteByName(t.Proposer)
	return err
}
