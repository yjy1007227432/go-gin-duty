package exchange_service

import (
	"go-gin-duty-master/models"
	"time"
)

type Exchange struct {
	Id            int       `form:"id"  json:"id"`
	RequestTime   string    `form:"request_time"  json:"request_time" binding:"required"`
	Proposer      string    `form:"proposer"   json:"proposer" binding:"required"`
	Respondent    string    `form:"respondent"   json:"respondent" binding:"required"`
	RequestedTime string    `form:"requested_time"  json:"requested_time"`
	Response      int       `form:"response"   json:"response"`
	ExchangeType  int       `form:"exchange_type"   json:"exchange_type" binding:"required"`
	CreatedOn     time.Time `form:"created_on"    json:"created_on"`
	ResponseOn    time.Time `form:"response_on"    json:"response_on"`
}

func (t *Exchange) GetAll() ([]models.DutyExchange, error) {
	exchanges, err := models.GetExchangeAll()
	return exchanges, err
}

func (t *Exchange) UpdateResponse() error {
	err := models.UpdateResponseExchange(t.Id, t.Response)
	return err
}

func (t *Exchange) GetExchangeByDate() ([]models.DutyExchange, error) {
	exchanges, err := models.GetExchangeByDate(t.RequestTime)
	return exchanges, err
}

func (t *Exchange) AddExchange() error {
	return models.AddExchange(t.RequestTime, t.Proposer, t.Respondent, t.RequestedTime, t.ExchangeType)
}

func (t *Exchange) IsExistDay() (bool, error) {
	isExist, err := models.IsExistDay(t.RequestTime, t.RequestedTime, t.Respondent, t.Proposer)
	return isExist, err
}

func (t *Exchange) DeleteAll() error {
	err := models.DeleteExchangeAll()
	return err
}

func (t *Exchange) GetMyExchange() ([]models.DutyExchange, error) {
	exchanges, err := models.GetMyExchange(t.Proposer, t.Response)
	return exchanges, err
}

func (t *Exchange) GetMyExamineExchange() ([]models.DutyExchange, error) {
	exchanges, err := models.GetMyExamineExchange(t.Respondent, t.Response)
	return exchanges, err
}

func (t *Exchange) GetExchangeById() (models.DutyExchange, error) {
	exchange, err := models.GetExchangeById(t.Id)
	return exchange, err
}

func (t *Exchange) DeleteById() error {
	err := models.DeleteExchangeById(t.Id)
	return err
}
func (t *Exchange) ExchangeTwo(group string) error {
	data := make(map[string]interface{})
	data["response"] = t.Response

	return models.ExchangeTwo(t.Id, group, data)
}
