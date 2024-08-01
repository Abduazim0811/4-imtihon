package rabbitmq

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func PublishOrder(body []byte) error {
	var conn *amqp.Connection
	var err error
	for i:=0; i<10;i++{
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			fmt.Println(err)
			time.Sleep(1*time.Second)
			continue
		}
	}
	if err != nil {
		return err
	}
	defer conn.Close()
	

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"command",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"command",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/grpc",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}