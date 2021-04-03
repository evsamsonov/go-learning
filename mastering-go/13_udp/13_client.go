package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")

		_, err := conn.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(text) == "exit" {
			fmt.Println("Exiting UDP client")
			return
		}

		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Reply: %s\n", string(buf[:n]))
	}
}
