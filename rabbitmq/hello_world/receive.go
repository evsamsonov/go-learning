package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
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
		log.Fatalf("failed to consume %v", err)
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received a messages: %s", d.Body)
		}
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(forever)
	}()

	log.Printf("[*] Waiting for messages. To exit press any key")
	<-forever
}
