package repositories

import (
	"errors"
	"kopikami/models"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	AddBatch(batch *models.Inventory) error
	GetInventoryByProduct(productID uint) ([]models.Inventory, error)
	UpdateStock(batchID uint, quantity int) error
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db}
}

func (r *inventoryRepository) AddBatch(batch *models.Inventory) error {
	var existingBatch models.Inventory
	if err := r.db.Where("batch_code = ?", batch.BatchCode).First(&existingBatch).Error; err == nil {
		return errors.New("batch code already exists")
	}
	return r.db.Create(batch).Error
}

func (r *inventoryRepository) GetInventoryByProduct(productID uint) ([]models.Inventory, error) {
	var inventories []models.Inventory
	err := r.db.Where("product_id = ?", productID).Order("expired_at ASC").Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepository) UpdateStock(batchID uint, quantity int) error {
	var inventory models.Inventory
	if err := r.db.First(&inventory, batchID).Error; err != nil {
		return err
	}
	if inventory.Quantity+quantity < 0 {
		return errors.New("stock cannot be negative")
	}
	return r.db.Model(&models.Inventory{}).Where("id = ?", batchID).Update("quantity", inventory.Quantity+quantity).Error
}
