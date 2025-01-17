package controllers

import (
	"kopikami/models"
	"kopikami/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RawMaterialBatchController struct {
    service services.RawMaterialBatchService
}

func NewRawMaterialBatchController(service services.RawMaterialBatchService) *RawMaterialBatchController {
    return &RawMaterialBatchController{service}
}

func (c *RawMaterialBatchController) Create(ctx *gin.Context) {
    var input struct {
        RawMaterialID  uint   `json:"raw_material_id"`
        BatchCode      string `json:"batch_code"`
        Quantity       int    `json:"quantity"`
        ReceivedDate   string `json:"received_date"`
        ExpirationDate string `json:"expiration_date"`
    }

    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ✅ Parsing tanggal dari string ke time.Time
    receivedDate, err := time.Parse("2006-01-02 15:04:05", input.ReceivedDate)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use 'YYYY-MM-DD HH:MM:SS'"})
        return
    }

    var expirationDate *time.Time
    if input.ExpirationDate != "" {
        parsedExpDate, err := time.Parse("2006-01-02 15:04:05", input.ExpirationDate)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration date format. Use 'YYYY-MM-DD HH:MM:SS'"})
            return
        }
        expirationDate = &parsedExpDate
    }

    batch := models.RawMaterialBatch{
        RawMaterialID:  input.RawMaterialID,
        BatchCode:      input.BatchCode,
        Quantity:       input.Quantity,
        ReceivedDate:   &receivedDate,
        ExpirationDate: expirationDate,
    }

    batchCreated, err := c.service.Create(batch)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, batchCreated)
}

func (c *RawMaterialBatchController) GetAll(ctx *gin.Context) {
    batches, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, batches)
}

func (c *RawMaterialBatchController) Delete(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
    if err := c.service.Delete(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Batch deleted successfully"})
}