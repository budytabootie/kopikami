package controllers

import (
    "kopikami/models"
    "kopikami/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type RawMaterialController struct {
    service services.RawMaterialService
}

func NewRawMaterialController(service services.RawMaterialService) *RawMaterialController {
    return &RawMaterialController{service}
}

// ✅ Menambahkan Konversi String ke Uint
func (c *RawMaterialController) Create(ctx *gin.Context) {
    var input models.RawMaterial
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    material, err := c.service.Create(input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, material)
}

func (c *RawMaterialController) GetAll(ctx *gin.Context) {
    materials, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, materials)
}

func (c *RawMaterialController) Update(ctx *gin.Context) {
    var input models.RawMaterial
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)  // ✅ Konversi string ke uint
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := c.service.Update(uint(id), input); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Material updated successfully"})
}

func (c *RawMaterialController) Delete(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)  // ✅ Konversi string ke uint
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := c.service.Delete(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}
