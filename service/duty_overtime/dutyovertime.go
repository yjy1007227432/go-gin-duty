package duty_overtime

import (
	"go-gin-duty-master/models"
	"time"
)

type DutyOverTime struct {
	Id         int
	Quantity   float64
	Proposer   string
	Reason     string `form:"reason"   json:"reason" `
	Checker    string
	Response   int
	CreatedOn  time.Time
	ResponseOn time.Time
}

func (t *DutyOverTime) GetAll() ([]models.DutyOvertime, error) {
	dutyOverTimes, err := models.GetOverTimeAll()
	return dutyOverTimes, err
}

func (t *DutyOverTime) GetAllNeedExamine() ([]models.DutyOvertime, error) {
	dutyOverTimes, err := models.GetOverTimeAllNeedExamine()
	return dutyOverTimes, err
}

func (t *DutyOverTime) AddOverTime() error {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["quantity"] = t.Quantity
	m["proposer"] = t.Proposer
	m["reason"] = t.Reason
	err := models.AddDutyOverTime(m)
	return err
}

func (t *DutyOverTime) GetOverTimeById() (models.DutyOvertime, error) {
	dutyOvertime, err := models.GetOverTimeById(t.Id)
	return dutyOvertime, err
}

func (t *DutyOverTime) Edit() error {
	data := make(map[string]interface{})
	data["response"] = t.Response

	return models.EditOverTime(t.Id, data)
}

func (t *DutyOverTime) GetOvertimesByName() ([]models.DutyOvertime, error) {
	dutyOvertimes, err := models.GetMyDutyOverTime(t.Proposer)
	return dutyOvertimes, err
}

func (t *DutyOverTime) DeleteOvertimeById() error {
	err := models.DeleteOvertimeById(t.Id)
	return err
}

func (t *DutyOverTime) AgreeAll() error {
	err := models.AgreeAll()
	return err
}

func (t *DutyOverTime) GetAllNowDay() ([]models.DutyOvertime, error) {
	overtimes, err := models.GetAllNowDay()
	return overtimes, err
}
