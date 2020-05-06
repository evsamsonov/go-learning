package server

import (
	"fmt"
	"github.com/evsamsonov/go-learning/microservices/1_rpc/contract"
	"log"
	"net"
	"net/rpc"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello, " + args.Name
	return nil
}

func StartServer() {
	err := rpc.Register(&HelloWorldHandler{})
	if err != nil {
		log.Fatalf("Failed to register: %v", err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen on port: %s", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Failed to accept: %s", err)
		}
		go rpc.ServeConn(conn)
	}
}
