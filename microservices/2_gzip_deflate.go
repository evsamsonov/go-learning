package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 8080

	http.Handle("/helloworld", http.HandlerFunc(helloWorldHandler))

	log.Printf("Server start on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorldHandler(rw http.ResponseWriter, _ *http.Request) {
	response := helloWorldResponse{
		Message: "Hello, world!",
	}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)

}
