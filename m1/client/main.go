package main

import (
	"context"
	"fmt"
	"log"

	example "example.com/networkpolicy/m1/server"

	"google.golang.org/grpc"
)

func runClient() error {
	conn, err := grpc.Dial("m2-service:50052", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	c := example.NewService2Client(conn)
	req := &example.GoodbyeRequest{Name: "John"}
	res, err := c.SayGoodbye(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println("m1 received:", res.Message)
	return nil
}

func main() {
	if err := runClient(); err != nil {
		log.Fatalf("Failed to start client: %v", err)
	}
}
