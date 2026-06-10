package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "server/proto/gen"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	addResp, err := client.Add(ctx, &pb.AddRequest{A: 10, B: 5})
	if err != nil {
		log.Fatal("Add failed: ", err)
	}
	log.Printf("Add(10, 5) = %d", addResp.Sum)

	subResp, err := client.Subtract(ctx, &pb.SubtractRequest{A: 10, B: 5})
	if err != nil {
		log.Fatal("Subtract failed: ", err)
	}
	log.Printf("Subtract(10, 5) = %d", subResp.Difference)

	mulResp, err := client.Multiply(ctx, &pb.MultiplyRequest{A: 10, B: 5})
	if err != nil {
		log.Fatal("Multiply failed: ", err)
	}
	log.Printf("Multiply(10, 5) = %d", mulResp.Product)

	divResp, err := client.Divide(ctx, &pb.DivideRequest{A: 10, B: 5})
	if err != nil {
		log.Fatal("Divide failed: ", err)
	}
	log.Printf("Divide(10, 5) = %f", divResp.Quotient)
}
