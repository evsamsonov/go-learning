package main

import (
	"context"
	pb "github.com/evsamsonov/go-learning/microservices/6_grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}()

	c := pb.NewGreaterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Ivan"})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("Greating: %s", r.GetMessage())
}
