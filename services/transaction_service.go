package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
)

type TransactionInput struct {
	UserID uint `json:"user_id" binding:"required"`
	Items  []TransactionItemInput `json:"items" binding:"required"`
}

type TransactionItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gte=1"`
}

type TransactionService interface {
	CreateTransaction(input TransactionInput) (*models.Transaction, error)
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionsByUserID(userID uint) ([]models.Transaction, error)
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
	productRepo     repositories.ProductRepository
	inventoryRepo   repositories.InventoryLogRepository
	productRecipeRepo repositories.ProductRecipeRepository
	rawMaterialRepo repositories.RawMaterialRepository
}

func NewTransactionService(
	transactionRepo repositories.TransactionRepository,
	productRepo repositories.ProductRepository,
	inventoryRepo repositories.InventoryLogRepository,
	productRecipeRepo repositories.ProductRecipeRepository,
	rawMaterialRepo repositories.RawMaterialRepository,
) TransactionService {
	return &transactionService{
		transactionRepo,
		productRepo,
		inventoryRepo,
		productRecipeRepo,
		rawMaterialRepo,
	}
}

func (s *transactionService) CreateTransaction(input TransactionInput) (*models.Transaction, error) {
	var totalAmount float64
	transaction := models.Transaction{
		UserID: input.UserID,
	}

	for _, item := range input.Items {
		product, err := s.productRepo.FindByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}

		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient product stock")
		}

		product.Stock -= item.Quantity
		err = s.productRepo.Update(&product)
		if err != nil {
			return nil, errors.New("failed to update product stock")
		}

		log := models.InventoryLog{
			Type:         "product",
			ReferenceID:  product.ID,
			ChangeAmount: -item.Quantity,
			Description:  "Stock reduction for transaction",
		}
		if err := s.inventoryRepo.Create(&log); err != nil {
			return nil, errors.New("failed to create inventory log for product")
		}

		transaction.Items = append(transaction.Items, models.TransactionItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		totalAmount += float64(item.Quantity) * product.Price

		recipes, err := s.productRecipeRepo.FindByProductID(item.ProductID)
		if err != nil {
			return nil, errors.New("failed to retrieve product recipes")
		}

		for _, recipe := range recipes {
			err := s.reduceRawMaterialStock(recipe.RawMaterialID, recipe.Quantity*item.Quantity)
			if err != nil {
				return nil, err
			}
		}
	}

	transaction.TotalAmount = totalAmount
	err := s.transactionRepo.Create(&transaction)
	return &transaction, err
}

func (s *transactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.transactionRepo.FindAll()
}

func (s *transactionService) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	return s.transactionRepo.FindByUserID(userID)
}

func (s *transactionService) reduceRawMaterialStock(rawMaterialID uint, requiredQuantity int) error {
	batches, err := s.inventoryRepo.GetBatchesByRawMaterialID(rawMaterialID)
	if err != nil {
		return errors.New("failed to retrieve raw material batches")
	}

	totalAvailable := 0
	for _, batch := range batches {
		totalAvailable += batch.Quantity
	}

	if totalAvailable < requiredQuantity {
		return errors.New("insufficient raw material stock")
	}

	remaining := requiredQuantity
	for _, batch := range batches {
		if remaining <= 0 {
			break
		}

		if batch.Quantity > remaining {
			batch.Quantity -= remaining
			remaining = 0
		} else {
			remaining -= batch.Quantity
			batch.Quantity = 0
		}

		if err := s.inventoryRepo.UpdateBatch(&batch); err != nil {
			return errors.New("failed to update raw material batch")
		}
	}

	log := models.InventoryLog{
		Type:         "raw_material",
		ReferenceID:  rawMaterialID,
		ChangeAmount: -requiredQuantity,
		Description:  "Stock reduction for transaction",
	}
	return s.inventoryRepo.Create(&log)
}
