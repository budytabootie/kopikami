package services

import (
	"kopikami/models"
	"kopikami/repositories"
	"time"
)

type DashboardService interface {
	GetSalesStats() (*models.SalesStats, error)
	GetInventoryStats() (*models.InventoryStats, error)
	GetSalesTrends(startDate, endDate time.Time) ([]models.SalesTrend, error)
}

type dashboardService struct {
	repo repositories.ReportRepository
}

func NewDashboardService(repo repositories.ReportRepository) DashboardService {
	return &dashboardService{repo}
}

func (s *dashboardService) GetSalesStats() (*models.SalesStats, error) {
	stats, err := s.repo.GetSalesStats()
	if err != nil {
		return nil, err
	}

	return &models.SalesStats{
		TotalRevenue:       stats.TotalRevenue,
		TotalTransactions:  stats.TotalTransactions,
		BestSellingProduct: stats.BestSellingProduct,
		QuantitySold:       stats.QuantitySold,
	}, nil
}

func (s *dashboardService) GetInventoryStats() (*models.InventoryStats, error) {
	stats, err := s.repo.GetInventoryStats()
	if err != nil {
		return nil, err
	}

	return &models.InventoryStats{
		ProductCount:     stats.ProductCount,
		RawMaterialCount: stats.RawMaterialCount,
	}, nil
}

func (s *dashboardService) GetSalesTrends(startDate, endDate time.Time) ([]models.SalesTrend, error) {
	// Ambil data dari repository
	trends, err := s.repo.GetSalesTrends(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Konversi data dari []repositories.SalesTrend ke []models.SalesTrend
	var result []models.SalesTrend
	for _, trend := range trends {
		result = append(result, models.SalesTrend{
			Date:       trend.Date,
			Revenue:    trend.Revenue,
		})
	}

	return result, nil
}

