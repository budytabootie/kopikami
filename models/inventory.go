package models

import (
    "time"
)

type Inventory struct {
    ItemID     int       `gorm:"primaryKey"`
    ItemName   string    `gorm:"size:255;not null"`
    Type       string    `gorm:"size:255;not null"`
    Stock      int       `gorm:"not null"`
    Price      int       `gorm:"not null"`
    CreatedAt  time.Time `gorm:"autoCreateTime"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime"`
    BatchID    int       `gorm:"not null"`
    Satuan     string    `gorm:"size:15;not null"`
    Notes      string    `gorm:"size:255"`
}
