package exchange_service

import (
	"go-gin-duty-master/models"
	"time"
)

type Exchange struct {
	Id            int
	RequestTime   string
	Proposer      string
	Respondent    string
	RequestedTime string
	Response      int
	CreatedOn     time.Time
	ResponseOn    time.Time
}

func (t *Exchange) GetAll() ([]models.DutyExchange, error) {
	exchanges, err := models.GetExchangeAll()
	return exchanges, err
}
