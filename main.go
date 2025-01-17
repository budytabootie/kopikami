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
	productRecipeRepo := repositories.NewProductRecipeRepository(db)
	inventoryRepo := repositories.NewInventoryLogRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	// Services
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo, inventoryRepo)
	rawMaterialService := services.NewRawMaterialService(rawMaterialRepo)
	rawMaterialBatchService := services.NewRawMaterialBatchService(rawMaterialBatchRepo, rawMaterialRepo, inventoryRepo)
	productRecipeService := services.NewProductRecipeService(productRecipeRepo, productRepo, rawMaterialRepo)
	inventoryService := services.NewInventoryService(inventoryRepo)
	transactionService := services.NewTransactionService(transactionRepo, productRepo, inventoryRepo, productRecipeRepo, rawMaterialRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)
	rawMaterialController := controllers.NewRawMaterialController(rawMaterialService)
	rawMaterialBatchController := controllers.NewRawMaterialBatchController(rawMaterialBatchService)
	productRecipeController := controllers.NewProductRecipeController(productRecipeService)
	inventoryController := controllers.NewInventoryController(inventoryService)
	transactionController := controllers.NewTransactionController(transactionService)

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

	// ✅ Product Recipe Routes
	protected.GET("/product-recipes", productRecipeController.GetAll)
	protected.GET("/product-recipes/:id", productRecipeController.GetByID)
	protected.POST("/product-recipes", productRecipeController.Create)
	protected.PUT("/product-recipes/:id", productRecipeController.Update)
	protected.DELETE("/product-recipes/:id", productRecipeController.Delete)

	// ✅ Inventory Log Routes
	protected.POST("/inventory", inventoryController.AddLog)
	protected.GET("/inventory", inventoryController.GetCurrentStock)

	// ✅ Transaction Routes
	protected.POST("/transactions", transactionController.CreateTransaction)
	protected.GET("/transactions", transactionController.GetAllTransactions)

	// Jalankan server di port 8080
	router.Run(":8080")
}
