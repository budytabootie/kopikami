package main

import (
	"net/http"
	"time"

	"kopikami/config"
	"kopikami/controllers"
	"kopikami/middlewares"
	"kopikami/repositories"
	"kopikami/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set timezone lokal
	time.Local = time.FixedZone("WIB", 7*3600) // Contoh untuk GMT+7 (WIB)

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
	reportRepo := repositories.NewReportRepository(db)

	// Services
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo, inventoryRepo)
	rawMaterialService := services.NewRawMaterialService(rawMaterialRepo)
	rawMaterialBatchService := services.NewRawMaterialBatchService(rawMaterialBatchRepo, rawMaterialRepo, inventoryRepo)
	productRecipeService := services.NewProductRecipeService(productRecipeRepo, productRepo, rawMaterialRepo)
	inventoryService := services.NewInventoryService(inventoryRepo)
	transactionService := services.NewTransactionService(transactionRepo, productRepo, inventoryRepo, productRecipeRepo, rawMaterialRepo)
	reportService := services.NewReportService(reportRepo)
	dashboardService := services.NewDashboardService(reportRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)
	rawMaterialController := controllers.NewRawMaterialController(rawMaterialService)
	rawMaterialBatchController := controllers.NewRawMaterialBatchController(rawMaterialBatchService)
	productRecipeController := controllers.NewProductRecipeController(productRecipeService)
	inventoryController := controllers.NewInventoryController(inventoryService)
	transactionController := controllers.NewTransactionController(transactionService)
	reportController := controllers.NewReportController(reportService)
	dashboardController := controllers.NewDashboardController(dashboardService)

	router := gin.Default()

	// ✅ Tambahkan Middleware CORS
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	})

	// ✅ Endpoint untuk Testing
	router.GET("/api/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ✅ Public Route (Tanpa Token)
	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/logout", authController.Logout) // Logout route

	router.POST("/api/inventory", inventoryController.AddLog)
	router.GET("/api/inventory", inventoryController.GetCurrentStock)

	// ✅ Protected Route (Menggunakan Token JWT)
	protected := router.Group("/api")
	protected.Use(middlewares.JWTMiddleware())

	// ✅ Routes untuk Role Admin
	adminRoutes := protected.Group("/admin") // Tambahkan prefiks /admin
	adminRoutes.Use(middlewares.RoleMiddleware("admin"))
	adminRoutes.GET("/products", productController.GetAllProducts)
	adminRoutes.POST("/products", productController.CreateProduct)
	adminRoutes.PUT("/products/:id", productController.UpdateProduct)
	adminRoutes.DELETE("/products/:id", productController.DeleteProduct)

	adminRoutes.POST("/raw-materials", rawMaterialController.Create)
	adminRoutes.GET("/raw-materials", rawMaterialController.GetAll)
	adminRoutes.PUT("/raw-materials/:id", rawMaterialController.Update)
	adminRoutes.DELETE("/raw-materials/:id", rawMaterialController.Delete)

	adminRoutes.POST("/raw-material-batches", rawMaterialBatchController.Create)
	adminRoutes.GET("/raw-material-batches", rawMaterialBatchController.GetAll)
	adminRoutes.DELETE("/raw-material-batches/:id", rawMaterialBatchController.Delete)

	adminRoutes.GET("/product-recipes", productRecipeController.GetAll)
	adminRoutes.GET("/product-recipes/:id", productRecipeController.GetByID)
	adminRoutes.POST("/product-recipes", productRecipeController.Create)
	adminRoutes.PUT("/product-recipes/:id", productRecipeController.Update)
	adminRoutes.DELETE("/product-recipes/:id", productRecipeController.Delete)

	adminRoutes.GET("/reports/sales", reportController.GetSalesReport)
	adminRoutes.GET("/reports/stock", reportController.GetStockReport)

	adminRoutes.GET("/dashboard/sales-stats", dashboardController.GetSalesStats)
	adminRoutes.GET("/dashboard/inventory-stats", dashboardController.GetInventoryStats)
	adminRoutes.GET("/dashboard/trends", dashboardController.GetSalesTrends)

	adminRoutes.GET("/transactions", transactionController.GetAllTransactions)

	// ✅ Routes untuk Role Kasir
	cashierRoutes := protected.Group("/cashier") // Tambahkan prefiks /cashier
	cashierRoutes.Use(middlewares.RoleMiddleware("cashier"))
	cashierRoutes.POST("/transactions", transactionController.CreateTransaction)
	cashierRoutes.GET("/transactions", func(ctx *gin.Context) {
		// Khusus kasir, ambil transaksi berdasarkan user_id
		userID, _ := ctx.Get("user_id")
		transactions, err := transactionController.GetTransactionsByUserID(userID.(uint))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, transactions)
	})

	// Jalankan server di port 8080
	router.Run(":8080")
}
