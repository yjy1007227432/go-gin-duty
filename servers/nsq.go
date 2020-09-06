package servers

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"go-gin-duty-master/e"
)

// 启动Nsq
func NsqRun() {
	Consumer()
}

// nsq发布消息
func Producer(msgBody string) {
	// 新建生产者
	p, err := nsq.NewProducer(e.HOST, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	// 发布消息
	if err := p.Publish(e.TOPIC_NAME, []byte(msgBody)); err != nil {
		panic(err)
	}
}

// nsq订阅消息
type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}

func Consumer() {
	c, err := nsq.NewConsumer(e.TOPIC_NAME, e.CHANNEL_NAME, nsq.NewConfig()) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	c.AddHandler(&ConsumerT{})                      // 添加消息处理
	if err := c.ConnectToNSQD(e.HOST); err != nil { // 建立连接
		panic(err)
	}
}
