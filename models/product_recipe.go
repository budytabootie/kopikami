package models

import "time"

type ProductRecipe struct {
	ID             uint      `gorm:"primaryKey"`
	ProductID      uint      `gorm:"not null"`
	RawMaterialID  uint      `gorm:"not null"`
	Quantity       int       `gorm:"not null;check:quantity > 0"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
