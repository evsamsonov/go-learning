package main

import (
	"context"
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
