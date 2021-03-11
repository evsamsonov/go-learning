package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main() {
	// any url
	req, err := http.NewRequest("GET", "http://localhost:6060/debug/pprof/", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", info)
		},
		PutIdleConn: nil,
		GotFirstResponseByte: func() {
			fmt.Println("Got first byte!")
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Done: %+v\n", info)
		},
		ConnectStart: func(_, _ string) {
			fmt.Println("Conn start")
		},
		ConnectDone: func(_, _ string, _ error) {
			fmt.Println("Conn done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), clientTrace))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, resp.Body)
}
