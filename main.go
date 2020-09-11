package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"go-gin-duty-master/pkg/setting"
	"go-gin-duty-master/routers"
	"go-gin-duty-master/service/timely_task"
	"log"
	"net/http"
)

func main() {
	//定时任务
	c := cron.New() // 新建一个定时任务对象
	c.AddFunc("0 0 23 * * ?", timely_task.ComputeVacation)
	c.AddFunc("0 30 8 * * ?", timely_task.AgreeMorningAndFullDay)
	c.AddFunc("0 0 14 * * ?", timely_task.AgreeAfternoon)

	c.AddFunc("0 30 8 * * ?", timely_task.AgreeDay)
	c.AddFunc("0 30 17 * * ?", timely_task.AgreeLate)
	c.Start()
	defer c.Stop()

	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter() //路由映射
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
