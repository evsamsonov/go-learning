package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"time"
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

	// prefetchCount = 1 - не позволяем консьюмеру брать больше,
	// чтобы освободившиеся консьюмеры могли сделать больше работы
	if ch.Qos(1, 0, false) != nil {
		log.Fatalf("failed to set prefetch")
	}

	var msgs <-chan amqp.Delivery
	msgs, err = ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("failed to consume")
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			time.Sleep(time.Duration(dotCount) * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [x] Waiting for messages. To exit press CTRL+C")
	<-forever
}
