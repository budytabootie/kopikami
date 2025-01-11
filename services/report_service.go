package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
)

type SalesReport struct {
	TotalSalesAmount float64 `json:"total_sales_amount"`
	TotalTransactions int    `json:"total_transactions"`
}

type ReportService interface {
	GenerateSalesReport() (*SalesReport, error)
}

type reportService struct {
	transactionRepo repositories.TransactionRepository
}

func NewReportService(transactionRepo repositories.TransactionRepository) ReportService {
	return &reportService{transactionRepo}
}

func (s *reportService) GenerateSalesReport() (*SalesReport, error) {
	transactions, err := s.transactionRepo.FindAll()
	if err != nil {
		return nil, errors.New("failed to fetch transactions")
	}

	var totalAmount float64
	totalTransactions := len(transactions)

	for _, t := range transactions {
		// Menggunakan struct models.Transaction secara eksplisit
		var transaction models.Transaction = t
		totalAmount += transaction.TotalAmount
	}

	return &SalesReport{
		TotalSalesAmount: totalAmount,
		TotalTransactions: totalTransactions,
	}, nil
}
