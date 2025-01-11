package services

import (
	"kopikami/models"
	"kopikami/repositories"
)

type ProductRecipeInput struct {
	ProductID     uint `json:"product_id"`
	RawMaterialID uint `json:"raw_material_id"`
	Quantity      int  `json:"quantity"`
}

type ProductRecipeService interface {
	CreateRecipe(input ProductRecipeInput) (*models.ProductRecipe, error)
}

type productRecipeService struct {
	recipeRepo repositories.ProductRecipeRepository
}

func NewProductRecipeService(recipeRepo repositories.ProductRecipeRepository) ProductRecipeService {
	return &productRecipeService{recipeRepo}
}

func (s *productRecipeService) CreateRecipe(input ProductRecipeInput) (*models.ProductRecipe, error) {
	recipe := models.ProductRecipe{
		ProductID:     input.ProductID,
		RawMaterialID: input.RawMaterialID,
		Quantity:      input.Quantity,
	}

	err := s.recipeRepo.Create(&recipe)
	return &recipe, err
}
