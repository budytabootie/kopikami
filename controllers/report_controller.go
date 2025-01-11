package controllers

import (
	"kopikami/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportService services.ReportService
}

func NewReportController(reportService services.ReportService) *ReportController {
	return &ReportController{reportService}
}

func (c *ReportController) GetSalesReport(ctx *gin.Context) {
	report, err := c.reportService.GenerateSalesReport()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, report)
}
