package models

import (
    "time"
)

type Stock struct {
    ItemID     int       `gorm:"primaryKey"`
    ItemName   string    `gorm:"size:255;not null"`
    Type       string    `gorm:"size:255;not null"`
    Stock      int       `gorm:"not null"`
    Price      int       `gorm:"not null"`
    BatchID    int       `gorm:"primaryKey"`
    Status     string    `gorm:"size:50;not null"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
