package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/adjust/rmq"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

const (
	maxSerialNumber = 1e6
)

type Message struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

type Queue interface {
	Add(messageName string, payload []byte) error
	AddMessage(message Message) error
	StartConsuming(prefetchLimit int64, pollDuration time.Duration, callback func(Message) error)
	Consume(delivery rmq.Delivery)
}

type RedisQueue struct {
	queue    rmq.Queue
	name     string
	callback func(Message) error
}

func (r *RedisQueue) Consume(delivery rmq.Delivery) {
	log.Printf("Got event from queue: %v", delivery.Payload())

	message := Message{}
	if err := json.Unmarshal([]byte(delivery.Payload()), &message); err != nil {
		log.Printf("Failed to parse event: %s\n", err.Error())
		delivery.Reject()
		return
	}

	if err := r.callback(message); err != nil {
		log.Printf("Failed to exec callback: %s\n", err.Error())
		delivery.Reject()
		return
	}

	delivery.Ack()
}

func NewRedisQueue(address, queueName string) (*RedisQueue, error) {
	errChan := make(chan error)
	conn, err := rmq.OpenConnection("my service", "tcp", address, 1, errChan)
	if err != nil {
		return nil, fmt.Errorf("open conn: %w", err)
	}
	queue, err := conn.OpenQueue(queueName)
	if err != nil {
		return nil, fmt.Errorf("open queue: %w", err)
	}
	return &RedisQueue{
		queue: queue,
		name:  queueName,
	}, nil
}

func (r *RedisQueue) Add(messageName string, payload []byte) error {
	message := Message{
		Name:    messageName,
		Payload: string(payload),
	}
	err := r.AddMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisQueue) AddMessage(message Message) error {
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(maxSerialNumber))
	if err != nil {
		return fmt.Errorf("rand int: %w", err)
	}
	message.ID = strconv.Itoa(time.Now().Nanosecond()) + serialNumber.String()
	payloadBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}
	err = r.queue.PublishBytes(payloadBytes)
	if err != nil {
		return fmt.Errorf("publish bytes: %w", err)
	}

	log.Printf("Added event to queue: %s\n", string(payloadBytes))
	return nil
}

func (r *RedisQueue) StartConsuming(prefetchLimit int64, pollDuration time.Duration, callback func(Message) error) {
	r.callback = callback
	r.queue.StartConsuming(prefetchLimit, pollDuration)
	r.queue.AddConsumer("RedisQueue_"+r.name, r)
}

func main() {
	// docker run --name some-redis -d redis
	queue, err := NewRedisQueue("localhost:6379", "test")
	if err != nil {
		log.Fatalf("Failed to new redis queue: %s\n", err.Error())
	}

	go func() {
		queue.StartConsuming(10, 100*time.Millisecond, func(message Message) error {
			log.Printf("Received message: %v\n", message)
			return nil
		})
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = queue.Add("new.product", data)
		if err != nil {
			log.Printf("Failed to add: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe("localhost:8080", nil)
}
