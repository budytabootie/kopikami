package repositories

import (
    "kopikami/models"
    "gorm.io/gorm"
)

type RawMaterialBatchRepository interface {
    Create(batch *models.RawMaterialBatch) error
    FindAll() ([]models.RawMaterialBatch, error)
    FindByID(id uint) (models.RawMaterialBatch, error)
    FindByRawMaterialID(rawMaterialID uint) ([]models.RawMaterialBatch, error)
    Update(batch *models.RawMaterialBatch) error
    Delete(id uint) error
}

type rawMaterialBatchRepository struct {
    db *gorm.DB
}

func NewRawMaterialBatchRepository(db *gorm.DB) RawMaterialBatchRepository {
    return &rawMaterialBatchRepository{db}
}

func (r *rawMaterialBatchRepository) Create(batch *models.RawMaterialBatch) error {
    return r.db.Create(batch).Error
}

func (r *rawMaterialBatchRepository) FindAll() ([]models.RawMaterialBatch, error) {
    var batches []models.RawMaterialBatch
    err := r.db.Find(&batches).Error
    return batches, err
}

func (r *rawMaterialBatchRepository) FindByID(id uint) (models.RawMaterialBatch, error) {
    var batch models.RawMaterialBatch
    err := r.db.First(&batch, id).Error
    return batch, err
}

func (r *rawMaterialBatchRepository) FindByRawMaterialID(rawMaterialID uint) ([]models.RawMaterialBatch, error) {
    var batches []models.RawMaterialBatch
    err := r.db.Where("raw_material_id = ?", rawMaterialID).Find(&batches).Error
    return batches, err
}

func (r *rawMaterialBatchRepository) Update(batch *models.RawMaterialBatch) error {
    return r.db.Save(batch).Error
}

func (r *rawMaterialBatchRepository) Delete(id uint) error {
    return r.db.Delete(&models.RawMaterialBatch{}, id).Error
}