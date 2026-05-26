// main.go
package main

import (
	"log"
	"net"

	pb "example.com/qabot/scheduler/proto"
	"example.com/qabot/scheduler/service"
	"google.golang.org/grpc"
)

func main() {
	// 1. Listen on a TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. Create a new gRPC server instance
	s := grpc.NewServer()

	// 3. Bind Service
	pb.RegisterHelloServer(s, &service.HelloServiceImpl{})
	log.Printf("Server listening on port :50051...")

	// 4. Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
