package client

import (
	"fmt"
	"github.com/evsamsonov/go-learning/microservices/1_rpc/contract"
	"log"
	"net/rpc"
)

const port = 1234

type Client struct {
	rpc *rpc.Client
}

func CreateClient() *Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &Client{rpc: client}
}

func (c *Client) PerformRequest() contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse
	err := c.rpc.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatalf("Failed to call rpc method: %v", err)
	}

	return reply
}
