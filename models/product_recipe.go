package models

type ProductRecipe struct {
	ID             uint   `gorm:"primaryKey"`
	ProductID      uint   `gorm:"not null"`
	RawMaterialID  uint   `gorm:"not null"`
	Quantity       int    `gorm:"not null"`
}
