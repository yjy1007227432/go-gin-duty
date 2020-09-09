package exchange_service

import (
	"go-gin-duty-master/models"
	"time"
)

type Exchange struct {
	Id            int       `json:"id"`
	RequestTime   string    `json:"request_time"`
	Proposer      string    `json:"proposer"`
	Respondent    string    `json:"respondent"`
	RequestedTime string    `json:"requested_time"`
	Response      int       `json:"response"`
	CreatedOn     time.Time `json:"created_on"`
	ResponseOn    time.Time `json:"response_on"`
}

func (t *Exchange) GetAll() ([]models.DutyExchange, error) {
	exchanges, err := models.GetExchangeAll()
	return exchanges, err
}

func (t *Exchange) AddExchange() error {

	return models.AddExchange(t.RequestTime, t.Proposer, t.Respondent, t.RequestedTime, 0)

}

func (t *Exchange) IsExistDay() (bool, error) {
	isExist, err := models.IsExistDay(t.RequestTime, t.RequestedTime)
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
func (t *Exchange) Edit() error {
	data := make(map[string]interface{})
	data["response"] = t.Response

	return models.EditExchange(t.Id, data)
}
