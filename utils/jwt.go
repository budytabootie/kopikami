package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secretKey adalah kunci rahasia untuk menandatangani token JWT
var secretKey = "supersecretkey"

// Claims mendefinisikan payload yang akan disimpan dalam token JWT
// Termasuk userID, role, dan informasi klaim standar

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT membuat token JWT baru berdasarkan userID dan role yang diberikan
// Token memiliki masa berlaku 24 jam
func GenerateJWT(userID uint, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Membuat token dengan metode HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateJWT memvalidasi token JWT yang diberikan
// Mengembalikan klaim jika token valid atau error jika token tidak valid
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("expired or invalid token")
	}

	return claims, nil
}
