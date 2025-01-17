package controllers

import (
	"fmt"
	"kopikami/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	service services.ReportService
}

func NewReportController(service services.ReportService) *ReportController {
	return &ReportController{service}
}

func (c *ReportController) GetSalesReport(ctx *gin.Context) {
	startDate, err := time.Parse("2006-01-02", ctx.Query("start_date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", ctx.Query("end_date"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
		return
	}

	// Debugging: Cetak nilai startDate dan endDate
	fmt.Printf("Start Date: %v, End Date: %v\n", startDate, endDate)

	report, err := c.service.GenerateSalesReport(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, report)
}


func (c *ReportController) GetStockReport(ctx *gin.Context) {
	report, err := c.service.GenerateStockReport()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, report)
}