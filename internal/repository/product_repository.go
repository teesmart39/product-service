package repository

import (
	"github.com/teesmart39/product-service/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetByID(id string) (*model.Product, error) {
	var product model.Product
	err := r.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (r *ProductRepository) List() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Find(&products).Error
	return products, err
}
