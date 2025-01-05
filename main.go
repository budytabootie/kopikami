package main

import (
	"kopikami/config"
	"kopikami/controllers"
	"kopikami/models"

	"github.com/gin-gonic/gin"
)

func main()  {
	config.InitDB()
	db := config.DB

	models.Migrate(db)

	r := gin.Default()

	r.GET("/inventories", controllers.GetInventory)

	r.Run(":8080")
}