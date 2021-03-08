package rabbitmq

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/streadway/amqp"
)

const PubName = "logs"

func Publish() {
	client := getClient()
	defer client.Close()

	ch, err := client.Channel()
	failOnError(err, "打开channel失败")
	defer ch.Close()

	stopChan := make(chan bool)
	go func() {
		pub(ch, stopChan)
	}()
	<-stopChan
}

// 发送到交换机
func pub(ch *amqp.Channel, c chan bool) {
	err := ch.ExchangeDeclare(
		PubName,  // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "声明一个queue失败")

	body := "num - "
	for i := 0; i < 10; i++ {
		err = ch.Publish(
			PubName, // exchange
			"",      // routing key
			false,   // mandatory
			false,   // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body + cast.ToString(i)),
			})

		fmt.Println("发送", body+cast.ToString(i))
		failOnError(err, "发送消息失败")

		time.Sleep(1 * time.Second)
	}
	time.Sleep(10 * time.Second)
	c <- true
}
