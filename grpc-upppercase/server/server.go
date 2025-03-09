package main

import (
	"context"
	"log"
	"net"
	"strings"

	pb "github.com/guijiangheng/uppercase/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUppercaseServiceServer
}

func (s *server) Uppercase(ctx context.Context, req *pb.UppercaseRequest) (*pb.UppercaseResponse, error) {
	output := strings.ToUpper(req.Input)
	return &pb.UppercaseResponse{Output: output}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUppercaseServiceServer(s, &server{})
	log.Printf("server listening on %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
