package models

import "time"

type RawMaterial struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);unique"`
	Unit      string    `gorm:"type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RawMaterialBatch struct {
	ID             uint       `gorm:"primaryKey"`
	RawMaterialID  uint       `gorm:"not null"`
	Quantity       int        `gorm:"not null"`
	ExpirationDate *time.Time
	CreatedAt      time.Time
}
