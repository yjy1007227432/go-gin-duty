package duty_overtime

import (
	"go-gin-duty-master/models"
	"time"
)

type DutyOverTime struct {
	Id         int
	Quantity   float64
	Proposer   string
	Reason     string
	Checker    string
	Response   int
	CreatedOn  time.Time
	ResponseOn time.Time
}

func (t *DutyOverTime) GetAll() ([]models.DutyOverTime, error) {
	dutyOverTimes, err := models.GetOverTimeAll()
	return dutyOverTimes, err
}

func (t *DutyOverTime) AddOverTime() error {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["quantity"] = t.Quantity
	m["proposer"] = t.Proposer
	m["reason"] = t.Reason
	err := models.AddDutyRest(m)
	return err
}

func (t *DutyOverTime) GetOverTimeById() (models.DutyOverTime, error) {
	dutyOvertime, err := models.GetOverTimeById(t.Id)
	return dutyOvertime, err
}

func (t *DutyOverTime) Edit() error {
	data := make(map[string]interface{})
	data["response"] = t.Response

	return models.EditOverTime(t.Id, data)
}

func (t *DutyOverTime) GetRestsByName() ([]models.DutyOverTime, error) {
	dutyOvertimes, err := models.GetMyDutyOverTime(t.Proposer)
	return dutyOvertimes, err
}
