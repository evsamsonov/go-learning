package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: trans}
	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
