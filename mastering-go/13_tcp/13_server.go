package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		req, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(req) == "exit" {
			return
		}

		fmt.Println("-> ", req)
		conn.Write([]byte(time.Now().Format(time.RFC3339) + "\n"))
	}
}
