package middlewares

import (
	"kopikami/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("Authorization")

		// Cek apakah token dikirim
		if tokenHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			ctx.Abort()
			return
		}

		// Validasi format "Bearer <token>"
		splitToken := strings.Split(tokenHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		// Validasi token JWT
		token := splitToken[1]
		if _, err := utils.ValidateJWT(token); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
