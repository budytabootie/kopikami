package controllers

import (
	"kopikami/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductRecipeController struct {
	productRecipeService services.ProductRecipeService
}

func NewProductRecipeController(productRecipeService services.ProductRecipeService) *ProductRecipeController {
	return &ProductRecipeController{productRecipeService}
}

func (c *ProductRecipeController) CreateRecipe(ctx *gin.Context) {
	var input services.ProductRecipeInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err := c.productRecipeService.CreateRecipe(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, recipe)
}
