package models

import "time"

type InventoryLog struct {
	ID            uint      `gorm:"primaryKey"`
	Type          string    `gorm:"type:enum('raw_material','product');not null"`
	ReferenceID   uint      `gorm:"not null"`
	ChangeAmount  int       `gorm:"not null"`
	Description   string    `gorm:"type:varchar(255)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
