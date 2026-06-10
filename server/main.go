package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "server/proto/gen"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(c context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Printf("Add: %d + %d = %d", req.A, req.B, sum)
	return &pb.AddResponse{Sum: sum}, nil
}

func (s *server) Subtract(c context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	difference := req.A - req.B
	log.Printf("Subtract: %d - %d = %d", req.A, req.B, difference)
	return &pb.SubtractResponse{Difference: difference}, nil
}

func (s *server) Multiply(c context.Context, req *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	product := req.A * req.B
	log.Printf("Multiply: %d * %d = %d", req.A, req.B, product)
	return &pb.MultiplyResponse{Product: product}, nil
}

func (s *server) Divide(c context.Context, req *pb.DivideRequest) (*pb.DivideResponse, error) {
	if req.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "division by zero is not allowed")
	}
	quotient := float32(req.A) / float32(req.B)
	log.Printf("Divide: %d / %d = %f", req.A, req.B, quotient)
	return &pb.DivideResponse{Quotient: quotient}, nil
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
