package main

import (
	"log"
	"net"

	"github.com/teesmart39/product-service/internal/handler"
	"github.com/teesmart39/product-service/internal/model"
	"github.com/teesmart39/product-service/internal/repository"
	"github.com/teesmart39/product-service/internal/service"
	productpb "github.com/teesmart39/product-service/proto"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/glebarez/sqlite"
)

func main() {
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.Product{})

	repo := repository.NewProductRepository(db)
	service := service.NewProductService(repo)

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServiceServer(grpcServer, handler.NewProductHandler(service))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
