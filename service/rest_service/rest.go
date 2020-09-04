package rest_service

import (
	"go-gin-duty-master/models"
	"time"
)

type Rest struct {
	Id         int
	Datetime   string
	Type       int
	Proposer   string
	Checker    string
	Response   int
	CreatedOn  time.Time
	ResponseOn time.Time
	Backup1    string
	Backup2    string
}

func (t *Rest) Add() error {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["type"] = t.Type
	m["request_time"] = t.Datetime
	m["proposer"] = t.Proposer
	m["checker"] = t.Checker
	m["response"] = t.Response
	err := models.AddDutyRest(m)
	return err
}
