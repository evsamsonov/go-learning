package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buf := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, _, err := conn.ReadFromUDP(buf)
		text := string(buf[:n-1])
		fmt.Println("-> ", text)
		if strings.TrimSpace(text) == "exit" {
			fmt.Println("Exiting UDP server")
			return
		}

		data := []byte(strconv.Itoa(rand.Intn(1000)))
		fmt.Printf("data: %s\n", string(data))
		n, err = conn.WriteToUDP(data, addr)
		if err != nil {
			log.Fatal(err)
			return
		}

		return
	}
}
