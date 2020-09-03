package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-gin-duty-master/pkg/setting"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	//设置表名为单数表名
	db.SingularTable(true)
}

func CloseDB() {
	defer db.Close()
}
