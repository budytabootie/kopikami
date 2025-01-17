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
    repo repositories.RawMaterialBatchRepository
}

func NewRawMaterialBatchService(repo repositories.RawMaterialBatchRepository) RawMaterialBatchService {
    return &rawMaterialBatchService{repo}
}

func (s *rawMaterialBatchService) Create(batch models.RawMaterialBatch) (*models.RawMaterialBatch, error) {
    if batch.Quantity <= 0 {
        return nil, errors.New("quantity must be greater than zero")
    }

    // ✅ Validasi hanya jika nilai tanggal benar-benar kosong
    if batch.ReceivedDate == nil {
        return nil, errors.New("received date is required")
    }

    err := s.repo.Create(&batch)
    return &batch, err
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
    _, err := s.GetByID(id)
    if err != nil {
        return errors.New("batch not found")
    }
    batch.ID = id
    return s.repo.Update(&batch)
}

func (s *rawMaterialBatchService) Delete(id uint) error {
    return s.repo.Delete(id)
}