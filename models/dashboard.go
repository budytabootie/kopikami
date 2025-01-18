package models

import "time"

type SalesStats struct {
	TotalRevenue       float64 `json:"total_revenue"`
	TotalTransactions  int     `json:"total_transactions"`
	BestSellingProduct string  `json:"best_selling_product"`
	QuantitySold       int     `json:"quantity_sold"`
}

type InventoryStats struct {
	ProductCount     int `json:"product_count"`
	RawMaterialCount int `json:"raw_material_count"`
}

type SalesTrend struct {
	Date    time.Time `json:"date"`
	Revenue float64   `json:"revenue"`
}
