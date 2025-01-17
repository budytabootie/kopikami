package repositories

import (
    "kopikami/models"
    "gorm.io/gorm"
)

type RawMaterialRepository interface {
    Create(material *models.RawMaterial) error
    FindAll() ([]models.RawMaterial, error)
    FindByID(id uint) (models.RawMaterial, error)
    Update(material *models.RawMaterial) error
    Delete(id uint) error
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

func (r *rawMaterialRepository) FindAll() ([]models.RawMaterial, error) {
    var materials []models.RawMaterial
    err := r.db.Find(&materials).Error
    return materials, err
}

func (r *rawMaterialRepository) FindByID(id uint) (models.RawMaterial, error) {
    var material models.RawMaterial
    err := r.db.First(&material, id).Error
    return material, err
}

func (r *rawMaterialRepository) Update(material *models.RawMaterial) error {
    return r.db.Save(material).Error
}

func (r *rawMaterialRepository) Delete(id uint) error {
    return r.db.Delete(&models.RawMaterial{}, id).Error
}