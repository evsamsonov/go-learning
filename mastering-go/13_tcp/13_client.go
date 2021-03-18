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
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Print(">> ")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(conn, text+"\n")
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("->: ", msg)

		if strings.TrimSpace(text) == "exit" {
			return
		}
	}
}
