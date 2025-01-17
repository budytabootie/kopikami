package repositories

import (
	"kopikami/models"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindAll() ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) FindAll() ([]models.Transaction, error) {
    var transactions []models.Transaction
    err := r.db.Preload("Items.Product").Find(&transactions).Error
    return transactions, err
}

func (r *transactionRepository) FindByDateRange(startDate, endDate time.Time) ([]models.Transaction, error) {
    var transactions []models.Transaction
    err := r.db.Preload("Items.Product").
        Where("created_at BETWEEN ? AND ?", startDate, endDate).
        Find(&transactions).Error
    return transactions, err
}
