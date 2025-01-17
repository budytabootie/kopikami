package models

import "time"

type RawMaterial struct {
    ID                uint      `gorm:"primaryKey"`
    Name              string    `gorm:"type:varchar(255);unique;not null" json:"name" binding:"required"`
    UnitOfMeasurement string    `gorm:"column:unit_of_measurement;type:varchar(50);not null" json:"unit_of_measurement" binding:"required"`
    Description       string    `gorm:"type:text" json:"description"`
    CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
