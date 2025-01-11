package repositories

import (
	"kopikami/models"
	"gorm.io/gorm"
)

type ProductRecipeRepository interface {
	Create(recipe *models.ProductRecipe) error
	GetByProductID(productID uint) ([]models.ProductRecipe, error)
}

type productRecipeRepository struct {
	db *gorm.DB
}

func NewProductRecipeRepository(db *gorm.DB) ProductRecipeRepository {
	return &productRecipeRepository{db}
}

func (r *productRecipeRepository) Create(recipe *models.ProductRecipe) error {
	return r.db.Create(recipe).Error
}

func (r *productRecipeRepository) GetByProductID(productID uint) ([]models.ProductRecipe, error) {
	var recipes []models.ProductRecipe
	err := r.db.Where("product_id = ?", productID).Find(&recipes).Error
	return recipes, err
}
