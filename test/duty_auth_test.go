package test

import (
	"fmt"
	"go-gin-duty-master/models"
	_ "go-gin-duty-master/pkg/setting"
)

func init() {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	fmt.Println(">>>>>>>>>>>>>>>>>")
	m["name"] = "姚俊毅"
	m["telephone"] = "18958049857"
	m["group"] = "crm"
	m["username"] = "yaojunyi"
	m["password"] = "1007227432"
	m["is_administrator"] = 0
	m["created_by"] = "姚俊毅"
	models.AddAuth(m)
}
