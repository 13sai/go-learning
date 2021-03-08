package rabbitmq

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/streadway/amqp"
)

const WorkName = "work"

func Work() {
	client := getClient()
	defer client.Close()

	ch, err := client.Channel()
	failOnError(err, "打开channel失败")
	defer ch.Close()

	err = ch.Qos(
		1,     //预取计数
		0,     //预取大小
		false, //全局
	)

	workTask(ch)
	go func() {
		workConsumer(ch, "001", true)
	}()
	go func() {
		workConsumer(ch, "002", false)
	}()
	go func() {
		workConsumer(ch, "003", true)
	}()
	time.Sleep(100 * time.Second)
	fmt.Println("over")
}

func workTask(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		WorkName, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "声明一个queue失败")

	body := "work任务来了 - "
	for i := 0; i < 10; i++ {
		err := ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType:  "text/plain",
				Body:         []byte(body + cast.ToString(i)),
				DeliveryMode: amqp.Persistent,
			},
		)

		fmt.Println("发送", body+cast.ToString(i))
		failOnError(err, "发送消息失败")
	}
}

func workConsumer(ch *amqp.Channel, i string, ack bool) {
	msgs, err := ch.Consume(
		WorkName, // queue
		"",       // consumer
		false,    // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)

	failOnError(err, "消费者注册失败")

	go func() {
		for d := range msgs {
			fmt.Println("收到消息", i, string(d.Body))
			if ack {
				d.Ack(false)
			} else {
				d.Nack(false, true)
			}

			time.Sleep(1 * time.Second)
		}
		fmt.Println("消息消费完了")
	}()
}
