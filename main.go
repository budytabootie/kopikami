package main

import (
	"kopikami/config"
	"kopikami/controllers"
	"kopikami/middlewares"
	"kopikami/repositories"
	"kopikami/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	rawMaterialRepo := repositories.NewRawMaterialRepository(db)
	rawMaterialBatchRepo := repositories.NewRawMaterialBatchRepository(db)

	// Services
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)
	rawMaterialService := services.NewRawMaterialService(rawMaterialRepo)
	rawMaterialBatchService := services.NewRawMaterialBatchService(rawMaterialBatchRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)
	rawMaterialController := controllers.NewRawMaterialController(rawMaterialService)
	rawMaterialBatchController := controllers.NewRawMaterialBatchController(rawMaterialBatchService)

	router := gin.Default()

	// ✅ Public Route (Tanpa Token)
	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/login", authController.Login)

	// ✅ Protected Route (Menggunakan Token JWT)
	protected := router.Group("/api")
	protected.Use(middlewares.JWTMiddleware())

	protected.GET("/products", productController.GetAllProducts)
	protected.POST("/products", productController.CreateProduct)
	protected.PUT("/products/:id", productController.UpdateProduct)
	protected.DELETE("/products/:id", productController.DeleteProduct)

	// ✅ Raw Material Routes
	protected.POST("/raw-materials", rawMaterialController.Create)
	protected.GET("/raw-materials", rawMaterialController.GetAll)
	protected.PUT("/raw-materials/:id", rawMaterialController.Update)
	protected.DELETE("/raw-materials/:id", rawMaterialController.Delete)

	// ✅ Raw Material Batch Routes
	protected.POST("/raw-material-batches", rawMaterialBatchController.Create)
	protected.GET("/raw-material-batches", rawMaterialBatchController.GetAll)
	protected.DELETE("/raw-material-batches/:id", rawMaterialBatchController.Delete)

	// Jalankan server di port 8080
	router.Run(":8080")
}
