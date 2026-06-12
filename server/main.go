package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "server/proto/gen"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(c context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B

	log.Println("Sum: ", sum)

	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &server{})

	log.Println("Server is running on port: ", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
