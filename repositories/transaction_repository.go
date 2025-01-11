package repositories

import (
	"kopikami/models"
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

// Implementasi Metode Create
func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

// Implementasi Metode FindAll
func (r *transactionRepository) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Items").Find(&transactions).Error
	return transactions, err
}
