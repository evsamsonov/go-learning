package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHelloWorld)

	http.ListenAndServe("localhost:8001", nil)
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, world!")
}
