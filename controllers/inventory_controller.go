package controllers

import (
	"kopikami/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	inventoryService services.InventoryService
}

func NewInventoryController(inventoryService services.InventoryService) *InventoryController {
	return &InventoryController{inventoryService}
}

func (c *InventoryController) AddInventory(ctx *gin.Context) {
	var input services.InventoryInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	inventory, err := c.inventoryService.AddInventory(input)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, inventory)
}

func (c *InventoryController) GetInventoryByProduct(ctx *gin.Context) {
	productID, _ := strconv.Atoi(ctx.Param("productID"))
	inventory, err := c.inventoryService.GetInventoryByProduct(uint(productID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No inventory found"})
		return
	}

	ctx.JSON(http.StatusOK, inventory)
}
