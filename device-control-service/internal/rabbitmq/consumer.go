package rabbitmq

import (
	"Device-service/internal/mongodb"
	"Device-service/pb/deviceproto"
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func ConsumeOrders() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(
		q.Name,
		"",
		"command",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var command *deviceproto.CommandRequest
			err := json.Unmarshal(d.Body, &command)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}

			mongodb, err := mongodb.NewDeviceMongoDb()
			if err != nil {
				log.Fatal(err)
			}
			_, err = mongodb.AddCommand(context.Background(), command)
			if err != nil {
				log.Printf("Error saving command to MongoDB: %s", err)
				continue
			}

			log.Printf("Buyurtma olindi va qayta ishlanmoqda: %+v", command)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}