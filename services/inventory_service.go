package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
	"time"
)

type InventoryInput struct {
	ProductID uint      `json:"product_id" binding:"required"`
	BatchCode string    `json:"batch_code" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,gte=0"`
	ExpiredAt *time.Time `json:"expired_at"`
}

type InventoryService interface {
	AddInventory(input InventoryInput) (*models.Inventory, error)
	GetInventoryByProduct(productID uint) ([]models.Inventory, error)
}

type inventoryService struct {
	inventoryRepo repositories.InventoryRepository
}

func NewInventoryService(inventoryRepo repositories.InventoryRepository) InventoryService {
	return &inventoryService{inventoryRepo}
}

func (s *inventoryService) AddInventory(input InventoryInput) (*models.Inventory, error) {
	if input.Quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}

	inventory := models.Inventory{
		ProductID: input.ProductID,
		BatchCode: input.BatchCode,
		Quantity:  input.Quantity,
		ExpiredAt: input.ExpiredAt,
	}

	err := s.inventoryRepo.AddBatch(&inventory)
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (s *inventoryService) GetInventoryByProduct(productID uint) ([]models.Inventory, error) {
	return s.inventoryRepo.GetInventoryByProduct(productID)
}
