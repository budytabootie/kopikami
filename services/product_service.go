package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
)

type ProductInput struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Stock int     `json:"stock" binding:"required,gte=0"`
}

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(input ProductInput) (*models.Product, error)
	UpdateProduct(id uint, input ProductInput) (*models.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo}
}

// ✅ Get All Products
func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.FindAll()
}

// ✅ Create Product
func (s *productService) CreateProduct(input ProductInput) (*models.Product, error) {
	if input.Stock < 0 {
		return nil, errors.New("stock cannot be negative")
	}

	product := models.Product{
		Name:  input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	if err := s.productRepo.Create(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

// ✅ Update Product
func (s *productService) UpdateProduct(id uint, input ProductInput) (*models.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock

	if err := s.productRepo.Update(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

// ✅ Delete Product
func (s *productService) DeleteProduct(id uint) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	return s.productRepo.Delete(&product)
}
