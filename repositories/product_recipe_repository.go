package repositories

import (
	"kopikami/models"
	"gorm.io/gorm"
)

type ProductRecipeRepository interface {
	Create(recipe *models.ProductRecipe) error
	FindAll() ([]models.ProductRecipe, error)
	FindByID(id uint) (models.ProductRecipe, error)
	Update(recipe *models.ProductRecipe) error
	Delete(id uint) error
	FindByProductID(productID uint) ([]models.ProductRecipe, error) // Tambahkan method ini
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

func (r *productRecipeRepository) FindAll() ([]models.ProductRecipe, error) {
	var recipes []models.ProductRecipe
	err := r.db.Find(&recipes).Error
	return recipes, err
}

func (r *productRecipeRepository) FindByID(id uint) (models.ProductRecipe, error) {
	var recipe models.ProductRecipe
	err := r.db.First(&recipe, id).Error
	return recipe, err
}

func (r *productRecipeRepository) Update(recipe *models.ProductRecipe) error {
	return r.db.Save(recipe).Error
}

func (r *productRecipeRepository) Delete(id uint) error {
	return r.db.Delete(&models.ProductRecipe{}, id).Error
}

// ✅ Implementasi FindByProductID
func (r *productRecipeRepository) FindByProductID(productID uint) ([]models.ProductRecipe, error) {
	var recipes []models.ProductRecipe
	err := r.db.Where("product_id = ?", productID).Find(&recipes).Error
	return recipes, err
}
