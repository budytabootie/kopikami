package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
	"time"
)

type SalesReport struct {
	TotalRevenue       float64 `json:"total_revenue"`
	TotalTransactions  int     `json:"total_transactions"`
	BestSellingProduct string  `json:"best_selling_product"`
	QuantitySold       int     `json:"quantity_sold"`
}

type StockReport struct {
	Products     []models.Product      `json:"products"`
	RawMaterials []RawMaterialStock    `json:"raw_materials"`
}

type RawMaterialStock struct {
	Name    string                  `json:"name"`
	Stock   int                     `json:"stock"`
	Batches []models.RawMaterialBatch `json:"batches"`
}

type ReportService interface {
	GenerateSalesReport(startDate, endDate time.Time) (*SalesReport, error)
	GenerateStockReport() (*StockReport, error)
}

type reportService struct {
	repo repositories.ReportRepository
}

func NewReportService(repo repositories.ReportRepository) ReportService {
	return &reportService{repo}
}

func (s *reportService) GenerateSalesReport(startDate, endDate time.Time) (*SalesReport, error) {
	transactions, err := s.repo.GetSalesReport(startDate, endDate)
	if err != nil {
		return nil, err
	}

	totalRevenue := 0.0
	totalTransactions := len(transactions)
	productSales := make(map[string]int)

	for _, transaction := range transactions {
		for _, item := range transaction.Items {
			if item.Product == nil {
				return nil, errors.New("product data is missing in transaction item")
			}
			productSales[item.Product.Name] += item.Quantity
			totalRevenue += float64(item.Quantity) * item.Price
		}
	}

	bestSellingProduct := ""
	quantitySold := 0
	for product, quantity := range productSales {
		if quantity > quantitySold {
			bestSellingProduct = product
			quantitySold = quantity
		}
	}

	return &SalesReport{
		TotalRevenue:       totalRevenue,
		TotalTransactions:  totalTransactions,
		BestSellingProduct: bestSellingProduct,
		QuantitySold:       quantitySold,
	}, nil
}

func (s *reportService) GenerateStockReport() (*StockReport, error) {
	products, rawMaterials, err := s.repo.GetStockReport()
	if err != nil {
		return nil, err
	}

	rawMaterialStocks := []RawMaterialStock{}
	for _, material := range rawMaterials {
		batches := []models.RawMaterialBatch{}
		stock := 0
		for _, batch := range material.Batches {
			stock += batch.Quantity
			// Set field IsNearExpiry
			batch.IsNearExpiry = batch.ExpirationDate != nil && batch.ExpirationDate.Before(time.Now().AddDate(0, 1, 0))
			batches = append(batches, batch)
		}
		rawMaterialStocks = append(rawMaterialStocks, RawMaterialStock{
			Name:    material.Name,
			Stock:   stock,
			Batches: batches,
		})
	}

	return &StockReport{
		Products:     products,
		RawMaterials: rawMaterialStocks,
	}, nil
}

