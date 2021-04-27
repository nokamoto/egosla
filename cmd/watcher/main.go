package main

import (
	"log"
	"net"

	pb "github.com/nokamoto/egosla/api"
	"google.golang.org/grpc"
)

func main() {
	port := ":9000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWatcherServiceServer(s, pb.UnimplementedWatcherServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
