package controllers

import (
	"kopikami/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) *TransactionController {
	return &TransactionController{transactionService}
}

// CreateTransaction handles creating a new transaction
func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var input services.TransactionInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := c.transactionService.CreateTransaction(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transaction)
}

// GetAllTransactions fetches all transactions
func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := c.transactionService.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}
