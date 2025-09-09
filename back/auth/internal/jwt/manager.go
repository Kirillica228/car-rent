package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecretkey") 

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
	Role int `json:"role"`
}

func GenerateToken(userID uint,role int) (string, error) {
	claims := Claims{
		UserID: userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // токен на 72 часа
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
