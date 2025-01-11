package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
)

type TransactionInput struct {
	UserID uint                   `json:"user_id" binding:"required"`
	Items  []TransactionItemInput `json:"items" binding:"required"`
}

type TransactionItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gte=1"`
}

type TransactionService interface {
	CreateTransaction(input TransactionInput) (*models.Transaction, error)
	GetAllTransactions() ([]models.Transaction, error)
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
	productRepo     repositories.ProductRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository, productRepo repositories.ProductRepository) TransactionService {
	return &transactionService{transactionRepo, productRepo}
}

func (s *transactionService) CreateTransaction(input TransactionInput) (*models.Transaction, error) {
	if input.UserID == 0 {
		return nil, errors.New("user ID cannot be zero")
	}

	transaction := models.Transaction{
		UserID: input.UserID,
	}

	if err := s.transactionRepo.Create(&transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (s *transactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.transactionRepo.FindAll()
}
