package models

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);unique;not null" json:"name" binding:"required"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price" binding:"required,gt=0"`
	Stock     int       `gorm:"not null;default:0" json:"stock" binding:"required,gte=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
