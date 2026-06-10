package main

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "server/proto/gen"
)

var s = &server{}

func TestAdd(t *testing.T) {
	resp, err := s.Add(context.Background(), &pb.AddRequest{A: 10, B: 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Sum != 15 {
		t.Errorf("expected 15, got %d", resp.Sum)
	}
}

func TestSubtract(t *testing.T) {
	resp, err := s.Subtract(context.Background(), &pb.SubtractRequest{A: 10, B: 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Difference != 5 {
		t.Errorf("expected 5, got %d", resp.Difference)
	}
}

func TestMultiply(t *testing.T) {
	resp, err := s.Multiply(context.Background(), &pb.MultiplyRequest{A: 4, B: 3})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Product != 12 {
		t.Errorf("expected 12, got %d", resp.Product)
	}
}

func TestDivide(t *testing.T) {
	resp, err := s.Divide(context.Background(), &pb.DivideRequest{A: 10, B: 4})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Quotient != 2.5 {
		t.Errorf("expected 2.5, got %f", resp.Quotient)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := s.Divide(context.Background(), &pb.DivideRequest{A: 10, B: 0})
	if err == nil {
		t.Fatal("expected error for division by zero, got nil")
	}
	st, ok := status.FromError(err)
	if !ok || st.Code() != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument, got %v", err)
	}
}
