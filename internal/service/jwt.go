package service

import (
	"eBlog/internal/configs"
	"eBlog/internal/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// GenerateToken генерирует JWT токен с кастомными полями
func (s *Service) GenerateToken(userID int, isRefreshToken bool) (string, error) {
	claims := models.CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Issuer: configs.AppSettings.AppParams.ServerName,
		},
	}

	if isRefreshToken == true {
		claims.IsRefreshToken = true
		claims.StandardClaims.ExpiresAt = int64(
			time.Duration(configs.AppSettings.AuthParams.RefreshTokenTtlDays) * 24 * time.Hour)
	} else {
		claims.IsRefreshToken = false
		claims.StandardClaims.ExpiresAt = int64(
			time.Duration(configs.AppSettings.AuthParams.AccessTokenTtlMinutes) * time.Minute)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

// ParseToken парсит JWT токен и возвращает кастомные поля
func (s *Service) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
