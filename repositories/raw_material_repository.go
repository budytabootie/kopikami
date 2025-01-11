package repositories

import (
	"kopikami/models"
	"gorm.io/gorm"
)

type RawMaterialRepository interface {
	Create(material *models.RawMaterial) error
	GetAll() ([]models.RawMaterial, error)
}

type rawMaterialRepository struct {
	db *gorm.DB
}

func NewRawMaterialRepository(db *gorm.DB) RawMaterialRepository {
	return &rawMaterialRepository{db}
}

func (r *rawMaterialRepository) Create(material *models.RawMaterial) error {
	return r.db.Create(material).Error
}

func (r *rawMaterialRepository) GetAll() ([]models.RawMaterial, error) {
	var materials []models.RawMaterial
	err := r.db.Find(&materials).Error
	return materials, err
}
