package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

var (
	//nsqd的地址，使用了tcp监听的端口
	tcpNsqdAddrr = "127.0.0.1:4160"
)

func main() {
	//初始化配置
	config := nsq.NewConfig()
	for i := 0; i < 100; i++ {
		//创建100个生产者
		tPro, err := nsq.NewProducer(tcpNsqdAddrr, config)
		if err != nil {
			fmt.Println(err)
		}
		//主题
		topic := "Insert"
		//主题内容
		tCommand := "new data!"
		//发布消息
		err = tPro.Publish(topic, []byte(tCommand))
		if err != nil {
			fmt.Println(err)
		}
	}

}
