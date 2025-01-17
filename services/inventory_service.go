package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
)

type InventoryInput struct {
	Type         string `json:"type" binding:"required,oneof=raw_material product"`
	ReferenceID  uint   `json:"reference_id" binding:"required"`
	ChangeAmount int    `json:"change_amount" binding:"required"`
	Description  string `json:"description"`
}

type InventoryService interface {
	AddLog(input InventoryInput) (*models.InventoryLog, error)
	GetCurrentStock(logType string, referenceID uint) (int, error)
}

type inventoryService struct {
	repo repositories.InventoryLogRepository
}

func NewInventoryService(repo repositories.InventoryLogRepository) InventoryService {
	return &inventoryService{repo}
}

func (s *inventoryService) AddLog(input InventoryInput) (*models.InventoryLog, error) {
	if input.ChangeAmount == 0 {
		return nil, errors.New("change amount cannot be zero")
	}

	log := models.InventoryLog{
		Type:         input.Type,
		ReferenceID:  input.ReferenceID,
		ChangeAmount: input.ChangeAmount,
		Description:  input.Description,
	}

	err := s.repo.Create(&log)
	return &log, err
}

func (s *inventoryService) GetCurrentStock(logType string, referenceID uint) (int, error) {
	return s.repo.GetCurrentStockByTypeAndID(logType, referenceID)
}
