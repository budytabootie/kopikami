package services

import (
    "errors"
    "kopikami/models"
    "kopikami/repositories"
)

type RawMaterialService interface {
    Create(material models.RawMaterial) (*models.RawMaterial, error)
    GetAll() ([]models.RawMaterial, error)
    GetByID(id uint) (*models.RawMaterial, error)
    Update(id uint, material models.RawMaterial) error
    Delete(id uint) error
}

type rawMaterialService struct {
    repo repositories.RawMaterialRepository
}

func NewRawMaterialService(repo repositories.RawMaterialRepository) RawMaterialService {
    return &rawMaterialService{repo}
}

func (s *rawMaterialService) Create(material models.RawMaterial) (*models.RawMaterial, error) {
    // ✅ Validasi data yang benar
    if material.Name == "" || material.UnitOfMeasurement == "" {
        return nil, errors.New("name and unit of measurement are required")
    }
    err := s.repo.Create(&material)
    return &material, err
}

func (s *rawMaterialService) GetAll() ([]models.RawMaterial, error) {
    return s.repo.FindAll()
}

func (s *rawMaterialService) GetByID(id uint) (*models.RawMaterial, error) {
    material, err := s.repo.FindByID(id)
    return &material, err
}

func (s *rawMaterialService) Update(id uint, material models.RawMaterial) error {
    existingMaterial, err := s.GetByID(id)
    if err != nil {
        return errors.New("material not found")
    }

    // ✅ Hindari overwriting `created_at`
    material.CreatedAt = existingMaterial.CreatedAt
    material.ID = id
    return s.repo.Update(&material)
}


func (s *rawMaterialService) Delete(id uint) error {
    return s.repo.Delete(id)
}
