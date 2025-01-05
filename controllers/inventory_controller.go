package controllers

import (
	"net/http"
	"kopikami/config"
	"kopikami/models"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context)  {
	var inventories []models.Inventory
	config.DB.Find(&inventories)
	c.JSON(http.StatusOK, inventories)
}