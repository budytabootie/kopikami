package services

import (
    "errors"
    "kopikami/models"
    "kopikami/repositories"
)

type RawMaterialBatchService interface {
    Create(batch models.RawMaterialBatch) (*models.RawMaterialBatch, error)
    GetAll() ([]models.RawMaterialBatch, error)
    GetByID(id uint) (*models.RawMaterialBatch, error)
    GetByRawMaterialID(rawMaterialID uint) ([]models.RawMaterialBatch, error)
    Update(id uint, batch models.RawMaterialBatch) error
    Delete(id uint) error
}

type rawMaterialBatchService struct {
    repo            repositories.RawMaterialBatchRepository
    rawMaterialRepo repositories.RawMaterialRepository
    inventoryRepo   repositories.InventoryLogRepository
}

func NewRawMaterialBatchService(
    repo repositories.RawMaterialBatchRepository,
    rawMaterialRepo repositories.RawMaterialRepository,
    inventoryRepo repositories.InventoryLogRepository,
) RawMaterialBatchService {
    return &rawMaterialBatchService{repo, rawMaterialRepo, inventoryRepo}
}

func (s *rawMaterialBatchService) Create(batch models.RawMaterialBatch) (*models.RawMaterialBatch, error) {
    // Validasi keberadaan Raw Material
    if _, err := s.rawMaterialRepo.FindByID(batch.RawMaterialID); err != nil {
        return nil, errors.New("raw material not found")
    }

    // Validasi Quantity
    if batch.Quantity <= 0 {
        return nil, errors.New("quantity must be greater than zero")
    }

    // Simpan batch ke database
    err := s.repo.Create(&batch)
    if err != nil {
        return nil, err
    }

    // Tambahkan log ke inventory_logs
    log := models.InventoryLog{
        Type:         "raw_material",
        ReferenceID:  batch.RawMaterialID,
        ChangeAmount: batch.Quantity,
        Description:  "Stock addition via batch creation",
    }

    if err := s.inventoryRepo.Create(&log); err != nil {
        return nil, errors.New("failed to create inventory log")
    }

    return &batch, nil
}

func (s *rawMaterialBatchService) GetAll() ([]models.RawMaterialBatch, error) {
    return s.repo.FindAll()
}

func (s *rawMaterialBatchService) GetByID(id uint) (*models.RawMaterialBatch, error) {
    batch, err := s.repo.FindByID(id)
    return &batch, err
}

func (s *rawMaterialBatchService) GetByRawMaterialID(rawMaterialID uint) ([]models.RawMaterialBatch, error) {
    return s.repo.FindByRawMaterialID(rawMaterialID)
}

func (s *rawMaterialBatchService) Update(id uint, batch models.RawMaterialBatch) error {
    existingBatch, err := s.GetByID(id)
    if err != nil {
        return errors.New("batch not found")
    }

    existingBatch.Quantity = batch.Quantity
    existingBatch.ReceivedDate = batch.ReceivedDate
    existingBatch.ExpirationDate = batch.ExpirationDate

    return s.repo.Update(existingBatch)
}

func (s *rawMaterialBatchService) Delete(id uint) error {
    // Validasi apakah batch ada
    _, err := s.GetByID(id)
    if err != nil {
        return errors.New("batch not found")
    }

    // Validasi apakah batch sedang digunakan
    inventoryLogs, err := s.repo.FindLogsByBatchID(id)
    if err != nil {
        return errors.New("failed to check batch usage")
    }
    if len(inventoryLogs) > 0 {
        return errors.New("cannot delete batch, it is currently in use")
    }

    return s.repo.Delete(id)
}
