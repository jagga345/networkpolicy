package main

import (
	"context"
	// "m-1/protos/github.com/yourusername/yourrepo/protos"

	// "networkpolicy/m1/server/protos/example.com/m1/server/protos"
	"fmt"
	"log"
	"net"

	greetpb "example.com/networkpolicy/m1/server/protos"

	// greetpb "example.com/networkpolicy/protos"
	// "example.com/networkpolicy/protos"
	// greetpb "protos"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedService1Server
}

func (s *server) SayHello(ctx context.Context, req *greetpb.HelloRequest) (*greetpb.HelloResponse, error) {
	fmt.Println("Received request:", req.Name)
	return &greetpb.HelloResponse{Message: "Hello from m1, " + req.Name}, nil
}

func runServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	greetpb.RegisterService1Server(s, &server{})
	fmt.Println("m1: Service1 is running on port 50051")
	return s.Serve(lis)
}

func main() {
	if err := runServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
