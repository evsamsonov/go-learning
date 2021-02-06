package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexcesaro/statsd"
)

type helloRequest struct {
	Name string `json:"name"`
}

func main() {
	statdClient, err := statsd.New()
	if err != nil {
		log.Fatalf("Failed to create statd client: %v", err)
	}
	defer statdClient.Close()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		defer statdClient.NewTiming().Send("hello.response_time")
		statdClient.Increment("hello.called")

		var hello helloRequest
		err := json.NewDecoder(r.Body).Decode(&hello)
		if err != nil {
			statdClient.Increment("hello.failed")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello, %s!\n", hello.Name)
		statdClient.Increment("hello.success")
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
