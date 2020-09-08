package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/pkg/setting"
	"go-gin-duty-master/routers"
	"log"
	"net/http"
)

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter() //路由映射
	//readTimeout := setting.ServerSetting.ReadTimeout
	//writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
		//ReadTimeout:    readTimeout,
		//WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
