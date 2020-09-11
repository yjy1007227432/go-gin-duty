package timely_task

import (
	"github.com/prometheus/common/log"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/service/duty_vacation"
	"go-gin-duty-master/service/exchange_service"
	"go-gin-duty-master/service/rest_service"
	"time"
)

//每天晚上23点跑定时任务
//todo 年休计算
func ComputeVacation() {
	nowDay := time.Now().Format("2006-01-02")

	rests, err := (&rest_service.Rest{
		Datetime: nowDay,
	}).GetRestByDayAgree()

	if err != nil {
		log.Errorf("ComputeVacation  GetRestByDay  run error: \v", err)
	}

	for _, rest := range rests {
		if rest.Type == 2 {
			err := (&duty_vacation.Vacation{
				Name: rest.Proposer,
			}).AddOne(rest.VacationType)
			if err != nil {
				log.Errorf("ComputeVacation AddOne run error: \v", err)
			}
		} else {
			err := (&duty_vacation.Vacation{
				Name: rest.Proposer,
			}).AddHalf(rest.VacationType)
			if err != nil {
				log.Errorf("ComputeVacation AddOne run error: \v", err)
			}
		}
	}
}

//8:30 同意当天上午的调休以及全天的调休
func AgreeMorningAndFullDay() {
	nowDay := time.Now().Format("2006-01-02")
	err := (&rest_service.Rest{
		Datetime: nowDay,
	}).AgreeMorningAndFullDay()

	if err != nil {
		log.Errorf("AgreeMorningAndFullDay run error: \v", err)
	}
}

//2:00 同意当天下午的调休
func AgreeAfternoon() {
	nowDay := time.Now().Format("2006-01-02")
	err := (&rest_service.Rest{
		Datetime: nowDay,
	}).AgreeAfternoon()

	if err != nil {
		log.Errorf("AgreeAfternoon run error: \v", err)
	}
}

//8:30 同意当天白天的换班以及全天以及特殊班的换班
func AgreeDay() {
	nowDay := time.Now().Format("2006-01-02")
	exchanges, err := (&exchange_service.Exchange{
		RequestTime: nowDay,
	}).GetExchangeByDate()

	for _, exchange := range exchanges {
		if exchange.ExchangeType != 1 {
			_, group, err := (&auth_service.Auth{
				Username: exchange.Proposer,
			}).GetNameByUsername()
			if err != nil {
				log.Errorf("AgreeDay run error: \v", err)
				return
			}
			err = (&exchange_service.Exchange{
				Id:       exchange.Id,
				Response: 1,
			}).ExchangeTwo(group)
		}
	}
	if err != nil {
		log.Errorf("AgreeDay run error: \v", err)
		return
	}
	return
}

//17:30 同意当天晚班换班
func AgreeLate() {
	nowDay := time.Now().Format("2006-01-02")
	exchanges, err := (&exchange_service.Exchange{
		RequestTime: nowDay,
	}).GetExchangeByDate()

	for _, exchange := range exchanges {
		if exchange.ExchangeType == 1 {
			_, group, err := (&auth_service.Auth{
				Username: exchange.Proposer,
			}).GetNameByUsername()
			if err != nil {
				log.Errorf("AgreeDay run error: \v", err)
				return
			}
			err = (&exchange_service.Exchange{
				Id:       exchange.Id,
				Response: 1,
			}).ExchangeTwo(group)
		}
	}
	if err != nil {
		log.Errorf("AgreeDay run error: \v", err)
		return
	}
	return
}