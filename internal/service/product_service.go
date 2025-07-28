package service

import (
	"github.com/google/uuid"
	"github.com/teesmart39/product-service/internal/model"
	"github.com/teesmart39/product-service/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) CreateProduct(name string, price float32, desc string) (*model.Product, error) {
	product := &model.Product{
		ID:          uuid.NewString(),
		Name:        name,
		Price:       price,
		Description: desc,
	}
	err := s.Repo.Create(product)
	return product, err
}

func (s *ProductService) GetProduct(id string) (*model.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) ListProducts() ([]model.Product, error) {
	return s.Repo.List()
}
