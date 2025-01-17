package repositories

import (
	"kopikami/models"
	"time"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetSalesReport(startDate, endDate time.Time) ([]models.Transaction, error)
	GetStockReport() (products []models.Product, rawMaterials []models.RawMaterial, err error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db}
}

func (r *reportRepository) GetSalesReport(startDate, endDate time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction

	// Menyesuaikan rentang waktu untuk query
	startDate = startDate.Truncate(24 * time.Hour)
	endDate = endDate.Truncate(24 * time.Hour).Add(24*time.Hour - time.Nanosecond)

	err := r.db.Preload("Items.Product").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&transactions).Error

	return transactions, err
}



func (r *reportRepository) GetStockReport() (products []models.Product, rawMaterials []models.RawMaterial, err error) {
	err = r.db.Find(&products).Error
	if err != nil {
		return
	}
	err = r.db.Preload("Batches").Find(&rawMaterials).Error
	return
}
