package controllers

import (
	"kopikami/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductRecipeController struct {
	service services.ProductRecipeService
}

func NewProductRecipeController(service services.ProductRecipeService) *ProductRecipeController {
	return &ProductRecipeController{service}
}

func (c *ProductRecipeController) GetAll(ctx *gin.Context) {
	recipes, err := c.service.GetAllRecipes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, recipes)
}

func (c *ProductRecipeController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	recipe, err := c.service.GetRecipeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, recipe)
}

func (c *ProductRecipeController) Create(ctx *gin.Context) {
	var input services.ProductRecipeInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err := c.service.CreateRecipe(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, recipe)
}

func (c *ProductRecipeController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	var input services.ProductRecipeInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err := c.service.UpdateRecipe(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, recipe)
}

func (c *ProductRecipeController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	if err := c.service.DeleteRecipe(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}
