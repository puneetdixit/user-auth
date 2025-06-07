package utils

import (
	"errors"
	"user-auth/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecret") // TODO: use env var in production

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.User) (string, error) {
	claims := Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*models.User, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	var user models.User
	if err := models.DB.First(&user, claims.UserID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
