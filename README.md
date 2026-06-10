# gRPC Calculator Practice

A simple gRPC calculator service in Go that supports Add, Subtract, Multiply, and Divide operations.

## Project Structure

```
grpc-practice/
├── server/          # gRPC server
│   ├── main.go
│   ├── main_test.go
│   └── proto/
│       ├── main.proto
│       └── gen/     # generated protobuf files
└── client/          # gRPC client
    └── main.go
```

## Prerequisites

- Go 1.22+
- protoc (Protocol Buffer compiler)
- protoc-gen-go and protoc-gen-go-grpc plugins

## Running the Server

```bash
cd server
go run main.go
```

The server listens on `:50051`.

## Running the Client

With the server running in a separate terminal:

```bash
cd client
go run main.go
```

## Running Tests

```bash
cd server
go test ./...
```

## Regenerating Proto Files

```bash
cd server
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/main.proto
cp proto/main.pb.go proto/gen/main.pb.go
cp proto/main_grpc.pb.go proto/gen/main_grpc.pb.go
rm proto/main.pb.go proto/main_grpc.pb.go
```

## Available RPCs

| RPC       | Request fields | Response field |
|-----------|---------------|---------------|
| Add       | A, B (int32)  | sum (int32)   |
| Subtract  | A, B (int32)  | difference (int32) |
| Multiply  | A, B (int32)  | product (int32) |
| Divide    | A, B (int32)  | quotient (float32) |

> Divide returns `InvalidArgument` error when B is 0.
