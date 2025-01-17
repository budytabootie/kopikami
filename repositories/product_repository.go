package repositories

import (
	"kopikami/models"
	"gorm.io/gorm"
)

// ProductRepository mendefinisikan interface untuk operasi data pada entitas Product
// Meliputi operasi dasar seperti Create, Read, Update, dan Delete (CRUD)
type ProductRepository interface {
	FindAll() ([]models.Product, error)                        // Mengambil semua produk dari database
	FindByID(id uint) (models.Product, error)                 // Mengambil produk berdasarkan ID
	Create(product *models.Product) error                     // Membuat produk baru di database
	Update(product *models.Product) error                     // Memperbarui data produk di database
	Delete(product *models.Product) error                     // Menghapus data produk dari database
}

// productRepository adalah implementasi dari ProductRepository
// Menggunakan GORM sebagai ORM untuk mengakses database
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository membuat instance baru dari productRepository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

// FindAll mengambil semua produk yang tersedia di database
func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

// FindByID mengambil data produk berdasarkan ID
func (r *productRepository) FindByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return product, err
}

// Create menambahkan data produk baru ke database
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// Update memperbarui data produk yang sudah ada di database
func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete menghapus data produk dari database
func (r *productRepository) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}
