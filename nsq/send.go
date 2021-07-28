package nsq

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/spf13/cast"
)

var producer *nsq.Producer

// 主函数
func send(topic string) {
	strIP1 := "127.0.0.1:4150"
	strIP2 := "127.0.0.1:4152"
	InitProducer(strIP1)
	for i := 0; i < 3; i++ {
		command := "msg-" + cast.ToString(i)
		for err := Publish(topic, command); err != nil; err = Publish(topic, command) {
			//切换IP重连
			strIP1, strIP2 = strIP2, strIP1
			InitProducer(strIP1)
		}
		time.Sleep(time.Second * 2)
	}

}

// 初始化生产者
func InitProducer(str string) {
	var err error
	fmt.Println("address: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
}

//发布消息
func Publish(topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic, []byte(message)) // 发布消息
		return err
	}
	return err
}
