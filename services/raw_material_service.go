package services

import (
	"kopikami/models"
	"kopikami/repositories"
)

type RawMaterialInput struct {
	Name string `json:"name" binding:"required"`
	Unit string `json:"unit" binding:"required"`
}

type RawMaterialService interface {
	CreateMaterial(input RawMaterialInput) (*models.RawMaterial, error)
	GetAllMaterials() ([]models.RawMaterial, error)
}

type rawMaterialService struct {
	materialRepo repositories.RawMaterialRepository
}

func NewRawMaterialService(materialRepo repositories.RawMaterialRepository) RawMaterialService {
	return &rawMaterialService{materialRepo}
}

func (s *rawMaterialService) CreateMaterial(input RawMaterialInput) (*models.RawMaterial, error) {
	material := models.RawMaterial{
		Name: input.Name,
		Unit: input.Unit,
	}

	err := s.materialRepo.Create(&material)
	if err != nil {
		return nil, err
	}

	return &material, nil
}

func (s *rawMaterialService) GetAllMaterials() ([]models.RawMaterial, error) {
	return s.materialRepo.GetAll()
}
