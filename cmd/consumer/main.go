package main

import (
	"poc-go-web-rabbit/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "PROCESS_EVENT")
	for msg := range msgs {
		println(string(msg.Body))
		msg.Ack(false)
	}
}
