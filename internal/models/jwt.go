package models

import "github.com/dgrijalva/jwt-go"

// CustomClaims определяет кастомные поля токена
type CustomClaims struct {
	UserID         int  `json:"user_id"`
	IsRefreshToken bool `json:"is_refresh_token"`
	jwt.StandardClaims
}
