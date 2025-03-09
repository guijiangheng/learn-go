package main

import (
	"context"
	"log"
	"time"

	pb "github.com/guijiangheng/uppercase/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewUppercaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Uppercase(ctx, &pb.UppercaseRequest{Input: "hello world"})
	if err != nil {
		log.Fatalf("could not uppercase: %v", err)
	}

	log.Printf("uppercase: %s", r.Output)
}
