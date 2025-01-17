package middlewares

import (
	"kopikami/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware adalah middleware yang memvalidasi token JWT untuk endpoint yang dilindungi
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Mengambil token dari header Authorization
		tokenHeader := ctx.GetHeader("Authorization")

		// Cek apakah token dikirim atau tidak
		if tokenHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			ctx.Abort()
			return
		}

		// Memisahkan token dengan format "Bearer <token>"
		splitToken := strings.Split(tokenHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		// Validasi token JWT menggunakan utilitas
		token := splitToken[1]
		if _, err := utils.ValidateJWT(token); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// Jika token valid, lanjutkan ke handler berikutnya
		ctx.Next()
	}
}
