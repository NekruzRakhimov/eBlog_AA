package contracts

import (
	"eBlog/internal/models"
)

type ServiceI interface {
	// User logic
	CreateUser(u models.User) error
	AuthenticateUser(username, password string) (accessToken string, refreshToken string, err error)
	GenerateHash(input string) (string, error)

	// Articles logic
	CreateArticle(a models.Article) error
	GetAllArticles(userID int, title string) ([]models.Article, error)
	GetArticleByID(id int) (models.Article, error)
	UpdateArticle(a models.Article) error
	DeleteArticle(id int) error

	// JWT logic
	GenerateToken(userID int, isRefreshToken bool) (string, error)
	ParseToken(tokenString string) (*models.CustomClaims, error)
}
