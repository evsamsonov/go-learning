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

	var q amqp.Queue
	q, err = ch.QueueDeclare(
		"", // автоматически сгенерированное имя
		false,
		false,
		true, // удалится после отключения
		false,
		nil,
	)

	if err != nil {
		log.Fatal("failed to declare queue")
	}

	// Прибиндим очередь
	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)

	if err != nil {
		log.Fatal("failed to bind queue")
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

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" Waiting fo logs...")
	<-forever
}
