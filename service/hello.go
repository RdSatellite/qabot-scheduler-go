// service/hello.go
package service

import (
	"context"
	"log"

	pb "example.com/qabot/scheduler/proto"
)

type HelloServiceImpl struct {
	pb.UnimplementedHelloServer
}

func (s *HelloServiceImpl) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received message from client: %s", req.GetMsg())

	return &pb.HelloResponse{
		Msg: "Hello, gRPC-go!",
	}, nil
}
