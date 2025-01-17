package controllers

import (
	"kopikami/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	service services.InventoryService
}

func NewInventoryController(service services.InventoryService) *InventoryController {
	return &InventoryController{service}
}

func (c *InventoryController) AddLog(ctx *gin.Context) {
	var input services.InventoryInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log, err := c.service.AddLog(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, log)
}

func (c *InventoryController) GetCurrentStock(ctx *gin.Context) {
	logType := ctx.Query("type")
	referenceIDStr := ctx.Query("reference_id")

	if logType == "" || referenceIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "type and reference_id are required"})
		return
	}

	// Konversi referenceID dari string ke uint
	referenceID, err := strconv.ParseUint(referenceIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid reference_id format"})
		return
	}

	stock, err := c.service.GetCurrentStock(logType, uint(referenceID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"stock": stock})
}
