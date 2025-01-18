package controllers

import (
	"fmt"
	"kopikami/models"
	"kopikami/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{service}
}

// CreateTransaction hanya dapat diakses oleh role "cashier"
func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != "cashier" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
		return
	}

	var input services.TransactionInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := c.service.CreateTransaction(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transaction)
}

// GetAllTransactions dapat diakses oleh role "admin" dan "cashier"
func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
    fmt.Println("GetAllTransactions endpoint hit") // Log untuk debugging
    transactions, err := c.service.GetAllTransactions()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, transactions)
}



// GetTransactionsByUserID digunakan untuk role "cashier" berdasarkan user_id
func (c *TransactionController) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
    return c.service.GetTransactionsByUserID(userID)
}
