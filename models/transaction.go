package models

import "time"

type Transaction struct {
	ID          uint              `gorm:"primaryKey"`
	UserID      uint              `gorm:"not null"`
	TotalAmount float64           `gorm:"not null"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime"`
	Items       []TransactionItem `gorm:"foreignKey:TransactionID"`
}

type TransactionItem struct {
	ID            uint    `gorm:"primaryKey"`
	TransactionID uint    `gorm:"not null"`
	ProductID     uint    `gorm:"not null"`
	Quantity      int     `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
