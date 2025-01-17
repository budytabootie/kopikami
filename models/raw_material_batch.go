package models

import "time"

type RawMaterialBatch struct {
	ID             uint       `gorm:"primaryKey"`
	RawMaterialID  uint       `gorm:"not null"`
	BatchCode      string     `gorm:"type:varchar(100);not null;unique"`
	Quantity       int        `gorm:"not null"`
	ReceivedDate   *time.Time `gorm:"not null"`
	ExpirationDate *time.Time
	Description    string     `gorm:"type:varchar(255)"` // Tambahkan field ini
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
}

