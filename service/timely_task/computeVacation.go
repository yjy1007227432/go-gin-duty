package timely_task

import (
	"github.com/prometheus/common/log"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/service/duty_vacation"
	"go-gin-duty-master/service/exchange_service"
	"go-gin-duty-master/service/rest_service"
	"go-gin-duty-master/service/rota_service"
	"strings"
	"time"
)

//每天晚上23点跑定时任务

func ComputeVacation() {
	nowDay := time.Now().Format("2006-01-02")

	rests, err := (&rest_service.Rest{
		Datetime: nowDay,
	}).GetRestByDayAgree()

	if err != nil {
		//	log.Errorf("ComputeVacation  GetRestByDay  run error: \v", err)
	}

	for _, rest := range rests {
		if rest.Type == 2 {
			err := (&duty_vacation.Vacation{
				Name: rest.Proposer,
			}).ReduceOne(rest.VacationType)
			if err != nil {
				log.Errorf("ComputeVacation AddOne run error: \v", err)
			}
		} else {
			err := (&duty_vacation.Vacation{
				Name: rest.Proposer,
			}).ReduceHalf(rest.VacationType)
			if err != nil {
				log.Errorf("ComputeVacation AddOne run error: \v", err)
			}
		}
	}
	//计算调休
	rota, err := (&rota_service.Rota{
		Datetime: nowDay,
	}).GetRotaByDay()
	if err != nil {
		log.Errorf("ComputeVacation  GetRotaByDay  run error: \v", err)
	}
	if rota.CrmDutySpecial == "周末" {
		err := (&duty_vacation.Vacation{
			Name: rota.BillingLate,
		}).AddOneHalf()
		if err != nil {
			log.Errorf("ComputeVacation AddOneHalf run error: \v", err)
		}
		err = (&duty_vacation.Vacation{
			Name: rota.CrmLate,
		}).AddHalf()
		if err != nil {
			log.Errorf("ComputeVacation AddOneHalf run error: \v", err)
		}
		dutyNames := strings.Split(rota.CrmWeekendDay, "、")
		for _, name := range dutyNames {
			err = (&duty_vacation.Vacation{
				Name: name,
			}).AddOne()
			if err != nil {
				log.Errorf("ComputeVacation AddOne run error: \v", err)
			}
		}
	} else {
		err := (&duty_vacation.Vacation{
			Name: rota.BillingLate,
		}).AddHalf()
		if err != nil {
			log.Errorf("ComputeVacation AddHalf run error: \v", err)
		}
		err = (&duty_vacation.Vacation{
			Name: rota.CrmLate,
		}).AddHalf()
		if err != nil {
			log.Errorf("ComputeVacation AddHalf run error: \v", err)
		}
	}

}

//8:30 同意所有调休
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
		//	log.Errorf("AgreeAfternoon run error: \v", err)
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
				//		log.Errorf("AgreeDay run error: \v", err)
				return
			}
			err = (&exchange_service.Exchange{
				Id:       exchange.Id,
				Response: 2,
			}).ExchangeTwo(group)
		}
	}
	if err != nil {
		//	log.Errorf("AgreeDay run error: \v", err)
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
		_, group, err := (&auth_service.Auth{
			Username: exchange.Proposer,
		}).GetNameByUsername()
		if err != nil {
			log.Errorf("AgreeDay run error: \v", err)
			return
		}
		err = (&exchange_service.Exchange{
			Id:       exchange.Id,
			Response: 2,
		}).ExchangeTwo(group)
	}
	if err != nil {
		log.Errorf("AgreeDay run error: \v", err)
		return
	}
	return
}
