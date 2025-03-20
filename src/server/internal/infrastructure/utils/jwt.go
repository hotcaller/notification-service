package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"service/internal/infrastructure/config"
	"time"
)

func GenerateJWTToken(userID int64) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Subject:   fmt.Sprintf("%d", userID),
	}
	cfg := config.GetJWT()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
