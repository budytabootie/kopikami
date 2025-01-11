package models

import "time"

type Transaction struct {
	ID            uint              `gorm:"primaryKey"`
	UserID        uint              `gorm:"not null"`
	TotalAmount   float64           `gorm:"type:decimal(10,2)"`
	TransactionAt time.Time
	Items         []TransactionItem `gorm:"foreignKey:TransactionID"`
}

type TransactionItem struct {
	ID            uint      `gorm:"primaryKey"`
	TransactionID uint      `gorm:"not null"`
	ProductID     uint      `gorm:"not null"`
	Quantity      int
	Price         float64
}
