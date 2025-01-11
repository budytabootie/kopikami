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

	// Services
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)

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

	// Jalankan server di port 8080
	router.Run(":8080")
}
