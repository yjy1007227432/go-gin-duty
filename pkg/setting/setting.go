package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var cfg *ini.File

func Setup() {
	var err error
	//先导入设置
	cfg, err = ini.Load("conf/app.ini")
	//导入设置失败报错，注意：使用log.Fatal 和 log.Panic 相关的函数时，会调用os.Exit(1)退出程序
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}


}