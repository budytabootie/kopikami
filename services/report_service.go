package services

import (
	"kopikami/models"
	"kopikami/repositories"
	"time"
)

type SalesReport struct {
	TotalRevenue       float64              `json:"total_revenue"`
	TotalTransactions  int                  `json:"total_transactions"`
	TopSellingProducts []TopSellingProduct `json:"top_selling_products"`
}

type TopSellingProduct struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Revenue   float64 `json:"revenue"`
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

// GenerateSalesReport creates a sales report based on the given date range
func (s *reportService) GenerateSalesReport(startDate, endDate time.Time) (*SalesReport, error) {
	// Fetch transactions
	transactions, err := s.repo.GetSalesReport(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Fetch top-selling products
	topProductsRepo, err := s.repo.GetTopSellingProducts(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Convert repository result to service struct
	topProducts := []TopSellingProduct{}
	for _, p := range topProductsRepo {
		topProducts = append(topProducts, TopSellingProduct{
			ProductID: p.ProductID,
			Name:      p.Name,
			Quantity:  p.Quantity,
			Revenue:   p.Revenue,
		})
	}

	// Calculate total revenue
	totalRevenue := 0.0
	for _, transaction := range transactions {
		totalRevenue += transaction.TotalAmount
	}

	return &SalesReport{
		TotalRevenue:       totalRevenue,
		TotalTransactions:  len(transactions),
		TopSellingProducts: topProducts,
	}, nil
}

// GenerateStockReport creates a stock report for all products and raw materials
func (s *reportService) GenerateStockReport() (*StockReport, error) {
	// Fetch product and raw material stock data
	products, rawMaterials, err := s.repo.GetStockReport()
	if err != nil {
		return nil, err
	}

	// Prepare raw material stock details
	rawMaterialStocks := []RawMaterialStock{}
	for _, material := range rawMaterials {
		batches := []models.RawMaterialBatch{}
		stock := 0

		for _, batch := range material.Batches {
			stock += batch.Quantity
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
