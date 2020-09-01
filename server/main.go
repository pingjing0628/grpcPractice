package main

import (
	"context"
	pb "github.com/pingjing0628/grpcPractice/proto"
	"log"
	"net"
	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedHelloServiceServer
}

func main() {
	addr := "127.0.0.1:9999"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server listening on", addr)

	gRPCServer := grpc.NewServer()

	// 將 service 註冊到 gRPC server 中
	pb.RegisterHelloServiceServer(gRPCServer, &service{})
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetGreeting())

	return &pb.HelloResponse{Reply: "Heelo, " + in.GetGreeting()}, nil
}