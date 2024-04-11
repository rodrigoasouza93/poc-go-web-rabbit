package main

import (
	"context"
	"poc-go-web-rabbit/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	ctx := context.Background()
	err = rabbitmq.Publish(ctx, ch, "Hello, World!", "EVENTS_NOTIFICATION")
	if err != nil {
		panic(err)
	}
}
