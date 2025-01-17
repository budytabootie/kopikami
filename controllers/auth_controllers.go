package controllers

import (
	"kopikami/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController mengatur endpoint autentikasi pengguna
// Menggunakan AuthService untuk memproses data

type AuthController struct {
	authService services.AuthService
}

// NewAuthController membuat instance baru dari AuthController
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

// Register memungkinkan pengguna untuk mendaftar akun baru
func (c *AuthController) Register(ctx *gin.Context) {
	var input services.RegisterInput
	// ✅ Validasi input dari JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Mendaftarkan pengguna melalui AuthService
	user, err := c.authService.Register(input)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// ✅ Response sukses dengan data pengguna
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// Login memungkinkan pengguna untuk melakukan autentikasi
func (c *AuthController) Login(ctx *gin.Context) {
	var input services.LoginInput
	// ✅ Validasi input dari JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Proses login dan pengembalian token JWT
	token, err := c.authService.Login(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
