package handler

import (
	"context"

	"github.com/teesmart39/product-service/internal/service"
	productpb "github.com/teesmart39/product-service/proto"
)

// ProductHandler implements the gRPC server interface defined in productpb
type ProductHandler struct {
	productpb.UnimplementedProductServiceServer
	Service *service.ProductService
}

// NewProductHandler initializes a new handler with the service injected
func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: svc}
}

// CreateProduct handles the gRPC call to create a new product
func (h *ProductHandler) CreateProduct(ctx context.Context, req *productpb.ProductRequest) (*productpb.ProductResponse, error) {
	product, err := h.Service.CreateProduct(req.Name, req.Price, req.Description)
	if err != nil {
		return nil, err
	}
	return &productpb.ProductResponse{
		Id:      product.ID,
		Message: "Product created successfully",
	}, nil
}

// GetProduct handles the gRPC call to retrieve a product by ID
func (h *ProductHandler) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.Product, error) {
	product, err := h.Service.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}
	return &productpb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
	}, nil
}

// ListProducts returns all products from the database
func (h *ProductHandler) ListProducts(ctx context.Context, _ *productpb.Empty) (*productpb.ProductList, error) {
	products, err := h.Service.ListProducts()
	if err != nil {
		return nil, err
	}
	var protoProducts []*productpb.Product
	for _, p := range products {
		protoProducts = append(protoProducts, &productpb.Product{
			Id:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
		})
	}
	return &productpb.ProductList{Products: protoProducts}, nil
}
