package main

import (
	"fmt"
	"github.com/evsamsonov/go-learning/microservices/1_rpc/client"
	"github.com/evsamsonov/go-learning/microservices/1_rpc/server"
)

func main() {
	go server.StartServer()

	cl := client.CreateClient()
	fmt.Println(cl.PerformRequest())
}
