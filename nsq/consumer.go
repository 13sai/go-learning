package nsq

import (
	"context"
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

// 消费者
type Consumer struct{}

// 主函数
func Receive(ctx context.Context, cancel context.CancelFunc, topic string) {
	defer cancel()
	// address := "127.0.0.1:4161"

	channel := topic + "-channel"
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second * 2
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0) //屏蔽系统日志
	c.AddConcurrentHandlers(&Consumer{}, 3)

	//建立NSQLookupd连接
	// if err := c.ConnectToNSQLookupd(address); err != nil {
	// 	panic(err)
	// }

	//建立多个nsqd连接
	if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150"}); err != nil {
		panic(err)
	}
	<-ctx.Done()
	c.Stop()
	fmt.Println("consumer exit")
}

// 处理消息
func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}
