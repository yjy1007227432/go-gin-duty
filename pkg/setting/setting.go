package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var cfg *ini.File

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var DatabaseSetting = &Database{}
var ServerSetting = &Server{}

func init() {
	var err error
	//先导入设置
	cfg, err = ini.Load("D:\\code\\src\\go-gin-duty-master\\conf\\app.ini")
	//导入设置失败报错，注意：使用log.Fatal 和 log.Panic 相关的函数时，会调用os.Exit(1)退出程序
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v) //go-ini 中可以采用 MapTo 的方式来映射结构体,分别将section的模块参数映射到实体类中
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
