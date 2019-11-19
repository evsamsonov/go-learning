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

	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("failed to declare exchange")
	}

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		log.Fatal("failed to publish")
	}

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	if len(args) < 2 || os.Args[1] == "" {
		return "hello"
	}

	return strings.Join(args[1:], " ")
}
