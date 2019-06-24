package main

import "sync"

// Компилировать вместе с 1_4_lissajous.go, предварительно закомментировав main() в 1_4_lissajous.go
//  go build 1_4_lissajous.go 1_7_webserver2.go

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", urlHandler)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/lissajous", func(writer http.ResponseWriter, request *http.Request) {
		lissajous(writer)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func urlHandler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

func countHandler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n", count)
	mu.Unlock()
}


