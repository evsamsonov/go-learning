package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@0.0.0.0:5672/")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare queue %v", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello, world!"),
		})

	if err != nil {
		log.Fatalf("failed to publish %v", err)
	}
}
