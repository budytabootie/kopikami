package repositories

import (
	"fmt"
	"kopikami/models"
	"time"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetSalesReport(startDate, endDate time.Time) ([]models.Transaction, error)
	GetStockReport() (products []models.Product, rawMaterials []models.RawMaterial, err error)
	GetTopSellingProducts(startDate, endDate time.Time) ([]TopSellingProduct, error)
	GetLowStockProducts(threshold int) ([]models.Product, error)
	GetExpiringRawMaterials(days int) ([]models.RawMaterialBatch, error)
	GetSalesStats() (*SalesStats, error)
	GetInventoryStats() (*InventoryStats, error)
	GetSalesTrends(startDate, endDate time.Time) ([]SalesTrend, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db}
}

// GetSalesReport retrieves transactions within a specific date range
func (r *reportRepository) GetSalesReport(startDate, endDate time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	startDate = startDate.Truncate(24 * time.Hour)
	endDate = endDate.Truncate(24 * time.Hour).Add(24*time.Hour - time.Nanosecond)

	err := r.db.Preload("Items.Product").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&transactions).Error

	return transactions, err
}

// GetStockReport retrieves the current stock of products and raw materials
func (r *reportRepository) GetStockReport() (products []models.Product, rawMaterials []models.RawMaterial, err error) {
	err = r.db.Find(&products).Error
	if err != nil {
		return
	}
	err = r.db.Preload("Batches").Find(&rawMaterials).Error
	return
}

// GetTopSellingProducts retrieves the top-selling products in a specific date range
func (r *reportRepository) GetTopSellingProducts(startDate, endDate time.Time) ([]TopSellingProduct, error) {
	type result struct {
		ProductID uint    `json:"product_id"`
		Name      string  `json:"name"`
		Quantity  int     `json:"quantity"`
		Revenue   float64 `json:"revenue"`
	}
	var topProducts []result

	startDate = startDate.Truncate(24 * time.Hour)
	endDate = endDate.Truncate(24 * time.Hour).Add(24*time.Hour - time.Nanosecond)

	err := r.db.Table("transaction_items").
		Select("products.id AS product_id, products.name AS name, SUM(transaction_items.quantity) AS quantity, SUM(transaction_items.price * transaction_items.quantity) AS revenue").
		Joins("JOIN products ON transaction_items.product_id = products.id").
		Joins("JOIN transactions ON transaction_items.transaction_id = transactions.id").
		Where("transactions.created_at BETWEEN ? AND ?", startDate, endDate).
		Group("products.id").
		Order("quantity DESC").
		Scan(&topProducts).Error

	resultTop := make([]TopSellingProduct, len(topProducts))
	for i, p := range topProducts {
		resultTop[i] = TopSellingProduct(p)
	}
	return resultTop, err
}

// GetLowStockProducts retrieves products with stock below a specific threshold
func (r *reportRepository) GetLowStockProducts(threshold int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("stock < ?", threshold).Find(&products).Error
	return products, err
}

// GetExpiringRawMaterials retrieves raw material batches expiring within a specified number of days
func (r *reportRepository) GetExpiringRawMaterials(days int) ([]models.RawMaterialBatch, error) {
	var batches []models.RawMaterialBatch
	cutoffDate := time.Now().AddDate(0, 0, days)

	err := r.db.Where("expiration_date IS NOT NULL AND expiration_date <= ?", cutoffDate).
		Find(&batches).Error

	return batches, err
}

// GetSalesStats retrieves aggregate sales statistics
func (r *reportRepository) GetSalesStats() (*SalesStats, error) {
	var stats SalesStats
	err := r.db.Raw(`
		SELECT 
			SUM(total_amount) AS total_revenue,
			COUNT(*) AS total_transactions,
			(
				SELECT p.name 
				FROM transaction_items ti
				JOIN products p ON ti.product_id = p.id
				GROUP BY p.id
				ORDER BY SUM(ti.quantity) DESC
				LIMIT 1
			) AS best_selling_product,
			(
				SELECT SUM(quantity)
				FROM transaction_items
				GROUP BY product_id
				ORDER BY SUM(quantity) DESC
				LIMIT 1
			) AS quantity_sold
		FROM transactions
	`).Scan(&stats).Error

	return &stats, err
}

// GetInventoryStats retrieves aggregate inventory statistics
func (r *reportRepository) GetInventoryStats() (*InventoryStats, error) {
	var stats InventoryStats
	err := r.db.Raw(`
		SELECT 
			(SELECT COUNT(*) FROM products) AS product_count,
			(SELECT COUNT(*) FROM raw_materials) AS raw_material_count
	`).Scan(&stats).Error

	return &stats, err
}

// GetSalesTrends retrieves revenue trends within a specific date range
func (r *reportRepository) GetSalesTrends(startDate, endDate time.Time) ([]SalesTrend, error) {
	var trends []SalesTrend

	// Tambahkan waktu default
	startDate = startDate.Truncate(24 * time.Hour)                                 // 00:00:00
	endDate = endDate.Truncate(24 * time.Hour).Add(24*time.Hour - time.Nanosecond) // Akhir hari

	fmt.Printf("Executing query with startDate: %v, endDate: %v\n", startDate, endDate)
	err := r.db.Raw(`
    SELECT DATE(t.created_at) AS date,
           SUM(t.total_amount) AS revenue
    FROM transactions t
    WHERE t.created_at BETWEEN ? AND ?
    GROUP BY DATE(t.created_at)
    ORDER BY DATE(t.created_at)
`, startDate, endDate).Scan(&trends).Error
	fmt.Printf("Query result: %v, Error: %v\n", trends, err)

	return trends, err
}

type TopSellingProduct struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Revenue   float64 `json:"revenue"`
}

type SalesStats struct {
	TotalRevenue       float64
	TotalTransactions  int
	BestSellingProduct string
	QuantitySold       int
}

type InventoryStats struct {
	ProductCount     int
	RawMaterialCount int
}

type SalesTrend struct {
	Date    time.Time
	Revenue float64
}
