# Product Microservice – Lithium Digital Backend Assessment

This is a gRPC-based Product Microservice built in Go, using GORM, SQLite, and Clean Architecture. It provides CRUD operations for Products and Subscription Plans, structured to showcase clear separation of concerns.

---

## Features

- Clean architecture (handlers, services, repositories)
- SQLite database using a pure Go driver
- gRPC endpoints for CRUD operations
- GORM-based models and migrations
- Ready for testing and extension
- Dockerizable structure (optional enhancement)

---

## Project Structure

```
product-service/
├── cmd/
│   └── server/              # Entry point (main.go)
├── internal/
│   ├── handler/             # gRPC handlers
│   ├── model/               # GORM models
│   ├── repository/          # Database access layer
│   └── service/             # Business logic
├── proto/                   # .proto definitions and generated Go code
├── go.mod / go.sum
└── README.md
```

---

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/YOUR_USERNAME/product-service.git
cd product-service
```

### 2. Run the gRPC Server

```bash
cd cmd/server
go run main.go
```

You should see:

```
gRPC server running on port 50051...
```

This starts the gRPC server and initializes the products.db file automatically.

---

## How to Test with Postman (gRPC)

1. Open Postman
2. Go to New > gRPC Request
3. Select your `proto/product.proto` file
4. Set server to: `localhost:50051`
5. Choose a method (e.g. `CreateProduct`)
6. Use the following sample payload:

```json
{
  "name": "Standard Plan",
  "price": 2999,
  "description": "Basic product plan"
}
```

---

## Assumptions & Design Decisions

- Used `modernc.org/sqlite` to avoid system-level C dependency (gcc)
- Focused on Products service only to meet MVP scope
- Subscription Plan structure is ready but may be expanded in a separate microservice for scaling
- gRPC was chosen as required for speed and binary communication
- Followed clean separation for testability and future maintainability

---

## Optional Enhancements (Not Yet Implemented)

- Dockerfile and docker-compose for local dev
- Authentication (JWT or API key)
- Unit and integration tests with mocks
- Pagination and filtering for list endpoints
- REST proxy (grpc-gateway) for external access

---

## Submission

Please clone this repo, push to GitHub, and send the public or private link (with access) to the Lithium Digital team before Monday, July 28, 2025, 12:00 PM.

---

Thank you for the opportunity.
