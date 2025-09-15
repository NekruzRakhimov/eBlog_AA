package controller

import (
	"eBlog/internal/models"
	"eBlog/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type SignUpRequest struct {
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignUp
// @Summary     Регистрация
// @Description Регистрирует нового пользователя
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request body SignUpRequest true "данные для регистрации"
// @Success     201 {object} CommonResponse
// @Failure     400 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, CommonError{err.Error()})
		return
	}

	if err := service.CreateUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{err.Error()})
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{"User created successfully!"})
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokensPairResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// SignIn
// @Summary     Вход
// @Description Вход в свой аккаунт
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request body SignInRequest true "данные для входа в аккаунт"
// @Success     200 {object} TokensPairResponse
// @Failure     400 {object} CommonError
// @Failure     500 {object} CommonError
// @Router      /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var input SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, CommonError{err.Error()})
		return
	}

	accessToken, refreshToken, err := service.AuthenticateUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommonError{err.Error()})
		return
	}

	c.JSON(http.StatusOK, TokensPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// RefreshTokenPair
// @Summary     Обновить токены
// @Description Обновить пару токенов (получить новый access и refresh)
// @Tags        Auth
// @Produce     json
// @Param       Refresh-Token header string true "Bearer Refresh Token"
// @Success     200 {object} TokensPairResponse
// @Failure     401 {object} CommonError
// @Router      /auth/refresh [get]
func RefreshTokenPair(c *gin.Context) {
	header := c.GetHeader("Refresh-Token")

	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	refreshToken := headerParts[1]

	claims, err := service.ParseToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if claims.IsRefreshToken != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "token is not refresh token",
		})
		return
	}

	accessToken, err := service.GenerateToken(claims.UserID, false)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	refreshToken, err = service.GenerateToken(claims.UserID, true)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
