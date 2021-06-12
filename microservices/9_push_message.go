package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
	// curl -H "Content-Type: application/json" -X POST -d  'test'  http://127.0.0.1:8080/insert
	conn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect: %s", err.Error())
	}
	defer conn.Close()

	go func() {
		conn.Subscribe("product.inserted", func(msg *nats.Msg) {
			fmt.Printf("Received message: %s %s", msg.Subject, string(msg.Data))
		})
	}()

	go func() {
		conn.Subscribe("product.inserted", func(msg *nats.Msg) {
			fmt.Printf("Received message: %s %s", msg.Subject, string(msg.Data))
		})
	}()

	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			log.Println("insert called")
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			conn.Publish("product.inserted", data)
		}
	})

	http.ListenAndServe("localhost:8080", nil)
}
