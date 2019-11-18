package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
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

	// durable = true, позволяет восстаналивать сообщения после рестарта rabbitmq
	// (!!!) при объявлении очереди, которая уже существует параметры не обновляются
	q, err := ch.QueueDeclare("tasks_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("failed to declare queue %v", err)
	}

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,		// Чтобы не терять
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)

	if err != nil {
		log.Fatalf("failed to publish message")
	}

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	if len(args) < 2 || os.Args[1] == "" {
		return "hello"
	}

	return strings.Join(args[1:], " ")
}
