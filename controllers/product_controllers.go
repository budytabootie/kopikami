package controllers

import (
	"kopikami/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController mengatur endpoint terkait produk
// Menggunakan service ProductService untuk pemrosesan data

type ProductController struct {
	productService services.ProductService
}

// NewProductController membuat instance baru ProductController
func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService}
}

// GetAllProducts mengambil semua produk yang tersedia
func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// CreateProduct menambahkan produk baru
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var input services.ProductInput
	// Validasi input JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Membuat produk dengan memanggil service
	product, err := c.productService.CreateProduct(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

// UpdateProduct memperbarui data produk yang sudah ada
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var input services.ProductInput
	// Mengonversi ID dari parameter URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Validasi input JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memperbarui produk melalui service
	product, err := c.productService.UpdateProduct(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct menghapus produk berdasarkan ID
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	// Mengonversi ID dari parameter URL
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Menghapus produk dengan memanggil service
	if err := c.productService.DeleteProduct(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
