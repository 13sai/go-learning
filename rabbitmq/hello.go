package rabbitmq

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cast"
	"github.com/streadway/amqp"
)

const QueueName = "hello"
const ConsumerName = "hello-consumer"

func getClient() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "连接失败")
	return conn
}

func Hello() {
	client := getClient()
	defer client.Close()

	ch, err := client.Channel()
	failOnError(err, "打开channel失败")
	defer ch.Close()

	stopChan := make(chan bool)
	go func() {
		producer(ch, stopChan)
	}()
	<-stopChan
	consumer(ch)
	<-stopChan
}

func producer(ch *amqp.Channel, c chan bool) {
	q, err := ch.QueueDeclare(
		QueueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "声明一个queue失败")

	body := "num - "
	for i := 0; i < 10; i++ {
		err := ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body + cast.ToString(i)),
			},
		)

		fmt.Println("发送", body+cast.ToString(i))
		failOnError(err, "发送消息失败")

		time.Sleep(1 * time.Second)
	}
	c <- true
}

func consumer(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		QueueName,    // queue
		ConsumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	failOnError(err, "消费者注册失败")

	go func() {
		for d := range msgs {
			fmt.Println("收到消息", string(d.Body))
		}
	}()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
