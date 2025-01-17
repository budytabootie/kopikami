package repositories

import (
	"kopikami/models"
	"gorm.io/gorm"
)

type InventoryLogRepository interface {
    Create(log *models.InventoryLog) error
    GetCurrentStockByTypeAndID(logType string, referenceID uint) (int, error)
    GetBatchesByRawMaterialID(rawMaterialID uint) ([]models.RawMaterialBatch, error)
    UpdateBatch(batch *models.RawMaterialBatch) error // Tambahkan method ini
}


type inventoryLogRepository struct {
	db *gorm.DB
}

func NewInventoryLogRepository(db *gorm.DB) InventoryLogRepository {
	return &inventoryLogRepository{db}
}

func (r *inventoryLogRepository) Create(log *models.InventoryLog) error {
	return r.db.Create(log).Error
}

func (r *inventoryLogRepository) FindAll() ([]models.InventoryLog, error) {
	var logs []models.InventoryLog
	err := r.db.Find(&logs).Error
	return logs, err
}

func (r *inventoryLogRepository) GetCurrentStockByTypeAndID(logType string, referenceID uint) (int, error) {
	var total int
	err := r.db.Model(&models.InventoryLog{}).
		Select("SUM(change_amount)").
		Where("type = ? AND reference_id = ?", logType, referenceID).
		Scan(&total).Error
	return total, err
}

func (r *inventoryLogRepository) GetBatchesByRawMaterialID(rawMaterialID uint) ([]models.RawMaterialBatch, error) {
    var batches []models.RawMaterialBatch
    err := r.db.Where("raw_material_id = ? AND quantity > 0", rawMaterialID).
        Order("received_date ASC").
        Find(&batches).Error
    return batches, err
}

func (r *inventoryLogRepository) UpdateBatch(batch *models.RawMaterialBatch) error {
    return r.db.Save(batch).Error
}
