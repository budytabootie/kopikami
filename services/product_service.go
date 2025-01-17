package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
)

// ProductInput mendefinisikan struktur data untuk input pembuatan atau pembaruan produk
type ProductInput struct {
	Name  string  `json:"name" binding:"required"`               // Nama produk yang wajib diisi
	Price float64 `json:"price" binding:"required"`             // Harga produk yang wajib diisi
	Stock int     `json:"stock" binding:"required,gte=0"`      // Stok produk yang wajib diisi dan tidak boleh negatif
}

// ProductService mendefinisikan kontrak untuk layanan yang berhubungan dengan produk
type ProductService interface {
	GetAllProducts() ([]models.Product, error)                      // Mengambil semua produk
	CreateProduct(input ProductInput) (*models.Product, error)     // Membuat produk baru
	UpdateProduct(id uint, input ProductInput) (*models.Product, error)  // Memperbarui data produk
	DeleteProduct(id uint) error                                   // Menghapus produk berdasarkan ID
}

// productService adalah implementasi dari ProductService yang menggunakan ProductRepository
// untuk melakukan operasi CRUD pada produk
type productService struct {
	productRepo   repositories.ProductRepository
	inventoryRepo repositories.InventoryLogRepository // Ditambahkan
}

// NewProductService membuat instance baru dari ProductService
func NewProductService(productRepo repositories.ProductRepository, inventoryRepo repositories.InventoryLogRepository) ProductService {
	return &productService{
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
	}
}

// GetAllProducts mengambil semua produk yang tersedia di database
func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.FindAll()
}

// CreateProduct menambahkan produk baru dengan validasi stok tidak negatif
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

    // ✅ Tambahkan log stok awal ke inventory_logs
    if input.Stock > 0 {
        log := models.InventoryLog{
            Type:         "product",
            ReferenceID:  product.ID,
            ChangeAmount: input.Stock,
            Description:  "Initial stock added via product creation",
        }
        if err := s.inventoryRepo.Create(&log); err != nil {
            return nil, errors.New("failed to create inventory log")
        }
    }

    return &product, nil
}


// UpdateProduct memperbarui data produk yang sudah ada berdasarkan ID
func (s *productService) UpdateProduct(id uint, input ProductInput) (*models.Product, error) {
    product, err := s.productRepo.FindByID(id)
    if err != nil {
        return nil, errors.New("product not found")
    }

    // Hitung perubahan stok
    stockChange := input.Stock - product.Stock

    // Perbarui data produk
    product.Name = input.Name
    product.Price = input.Price
    product.Stock = input.Stock

    if err := s.productRepo.Update(&product); err != nil {
        return nil, err
    }

    // Tambahkan log ke inventory_logs jika ada perubahan stok
    if stockChange != 0 {
        log := models.InventoryLog{
            Type:         "product",
            ReferenceID:  product.ID,
            ChangeAmount: stockChange,
            Description:  "Stock updated via product update",
        }
        if err := s.inventoryRepo.Create(&log); err != nil {
            return nil, errors.New("failed to create inventory log")
        }
    }

    return &product, nil
}



// DeleteProduct menghapus produk dari database berdasarkan ID
func (s *productService) DeleteProduct(id uint) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	return s.productRepo.Delete(&product)
}
