package models

import "time"

type Inventory struct {
	ID         uint       `gorm:"primaryKey"`
	ProductID  uint       `gorm:"not null" json:"product_id" binding:"required"`
	BatchCode  string     `gorm:"type:varchar(100);not null" json:"batch_code" binding:"required"`
	Quantity   int        `gorm:"not null;default:0" json:"quantity" binding:"required,gte=0"`
	CreatedAt  time.Time  `json:"created_at"`
	ExpiredAt  *time.Time `json:"expired_at"`
}
