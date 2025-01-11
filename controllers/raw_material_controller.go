package controllers

import (
	"kopikami/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RawMaterialController struct {
	rawMaterialService services.RawMaterialService
}

func NewRawMaterialController(rawMaterialService services.RawMaterialService) *RawMaterialController {
	return &RawMaterialController{rawMaterialService}
}

func (c *RawMaterialController) CreateRawMaterial(ctx *gin.Context) {
	var input services.RawMaterialInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	material, err := c.rawMaterialService.CreateMaterial(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, material)
}

func (c *RawMaterialController) GetAllRawMaterials(ctx *gin.Context) {
	materials, err := c.rawMaterialService.GetAllMaterials()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, materials)
}
