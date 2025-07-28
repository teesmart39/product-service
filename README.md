# Product Microservice (Go + gRPC)

## Features
- gRPC CRUD API for products
- SQLite database via GORM
- Clean architecture structure

## Setup

1. Install Go, protoc, and plugins
2. Generate protobuf:
```bash
protoc --go_out=. --go-grpc_out=. proto/product.proto
```

3. Run the server:
```bash
go run cmd/server/main.go
```

Server runs on `localhost:50051`
