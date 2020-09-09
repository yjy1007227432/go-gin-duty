package timely_task

import (
	"github.com/prometheus/common/log"
	"go-gin-duty-master/service/duty_vacation"
	"go-gin-duty-master/service/rest_service"
	"time"
)

//每天晚上23点跑定时任务
func ComputeVacation() {
	nowDay := time.Now().Format("2006-01-02")

	rests, err := (&rest_service.Rest{
		Datetime: nowDay,
	}).GetRestByDay()

	if err != nil {
		log.Errorf("ComputeVacation  GetRestByDay  run error: \v", err)
	}

	for _, rest := range rests {
		if rest.Type == 1 {
			err := duty_vacation.Vacation{
				Name: rest.Proposer,
			}.AddOne()
			if err != nil {
				log.Errorf("ComputeVacation AddOne run error: \v", err)
			}
		} else {
			err := duty_vacation.Vacation{
				Name: rest.Proposer,
			}.AddHalf()
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
