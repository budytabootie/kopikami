package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
)

type ProductRecipeInput struct {
	ProductID     uint `json:"product_id" binding:"required"`
	RawMaterialID uint `json:"raw_material_id" binding:"required"`
	Quantity      int  `json:"quantity" binding:"required,gt=0"`
}

type ProductRecipeService interface {
	GetAllRecipes() ([]models.ProductRecipe, error)
	GetRecipeByID(id uint) (*models.ProductRecipe, error)
	CreateRecipe(input ProductRecipeInput) (*models.ProductRecipe, error)
	UpdateRecipe(id uint, input ProductRecipeInput) (*models.ProductRecipe, error)
	DeleteRecipe(id uint) error
}

type productRecipeService struct {
	repo             repositories.ProductRecipeRepository
	productRepo      repositories.ProductRepository
	rawMaterialRepo  repositories.RawMaterialRepository
}

func NewProductRecipeService(
	repo repositories.ProductRecipeRepository,
	productRepo repositories.ProductRepository,
	rawMaterialRepo repositories.RawMaterialRepository,
) ProductRecipeService {
	return &productRecipeService{repo, productRepo, rawMaterialRepo}
}

func (s *productRecipeService) GetAllRecipes() ([]models.ProductRecipe, error) {
	return s.repo.FindAll()
}

func (s *productRecipeService) GetRecipeByID(id uint) (*models.ProductRecipe, error) {
	recipe, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("recipe not found")
	}
	return &recipe, nil
}

func (s *productRecipeService) CreateRecipe(input ProductRecipeInput) (*models.ProductRecipe, error) {
	// Validasi keberadaan Product
	product, err := s.productRepo.FindByID(input.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	// Validasi keberadaan RawMaterial
	rawMaterial, err := s.rawMaterialRepo.FindByID(input.RawMaterialID)
	if err != nil {
		return nil, errors.New("raw material not found")
	}

	// Membuat ProductRecipe
	recipe := models.ProductRecipe{
		ProductID:     product.ID,
		RawMaterialID: rawMaterial.ID,
		Quantity:      input.Quantity,
	}

	// Simpan ke database
	err = s.repo.Create(&recipe)
	return &recipe, err
}

func (s *productRecipeService) UpdateRecipe(id uint, input ProductRecipeInput) (*models.ProductRecipe, error) {
	recipe, err := s.GetRecipeByID(id)
	if err != nil {
		return nil, err
	}

	// Update data recipe
	recipe.ProductID = input.ProductID
	recipe.RawMaterialID = input.RawMaterialID
	recipe.Quantity = input.Quantity

	// Simpan perubahan
	err = s.repo.Update(recipe)
	return recipe, err
}

func (s *productRecipeService) DeleteRecipe(id uint) error {
	recipe, err := s.GetRecipeByID(id)
	if err != nil {
		return errors.New("recipe not found")
	}

	return s.repo.Delete(recipe.ID)
}
