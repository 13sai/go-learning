package nsq

import (
	"context"
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/spf13/cast"
)

// 主函数
func Send(ctx context.Context, cancel context.CancelFunc, topic string) {
	defer cancel()
	str := "127.0.0.1:4150"
	fmt.Println("address: ", str)
	producer, err := nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	producer.SetLogger(nil, 0)

	for i := 0; i < 5; i++ {
		msg := "13sai, " + cast.ToString(i)
		fmt.Println("publish", msg, producer.Publish(topic, []byte(msg)))
		time.Sleep(time.Second * 1)
	}

	<-ctx.Done()
	producer.Stop()
	fmt.Println("producer exit")
}
