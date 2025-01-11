package repositories

import (
	"kopikami/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id uint) (models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(product *models.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

// ✅ Find All Products
func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

// ✅ Find Product by ID
func (r *productRepository) FindByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return product, err
}

// ✅ Create Product
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// ✅ Update Product
func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// ✅ Delete Product
func (r *productRepository) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}
