package main

import (
	"context"
	"fmt"
	pb "github.com/evsamsonov/go-learning/microservices/6_grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreaterServer
}

func (s *server) SayHello(_ context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func (s *server) SayStreamHello(req *pb.HelloRequest, serv pb.Greater_SayStreamHelloServer) error {
	for _, sym := range req.GetName() {
		err := serv.Send(&pb.HelloReply{Message: "Hello " + string(sym)})
		if err != nil {
			return fmt.Errorf("failed to send: %w", err)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreaterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %s", err)
	}
}
