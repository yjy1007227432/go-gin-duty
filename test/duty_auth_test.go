package test

import (
	"fmt"
	"go-gin-duty-master/models"
	_ "go-gin-duty-master/pkg/setting"
	"log"
	"testing"
)

func TestAddAuth(t *testing.T) {
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

func TestUpdateAuth(t *testing.T) {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["id"] = 2
	m["name"] = "姚俊毅"
	m["telephone"] = "18958049857"
	m["group"] = "crm"
	m["username"] = "yaojunyi"
	m["password"] = "1007227432yjy"
	m["is_administrator"] = 0
	m["modified_by"] = "姚俊毅"
	err := models.UpdateAuth(m)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}

func TestDeleteAuthById(t *testing.T) {
	err := models.DeleteAuthById(2)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}

func TestCheckAuth(t *testing.T) {
	_, err := models.CheckAuth("yaojunyi", "1007227432yjy")
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
