package contracts

import "eBlog/internal/models"

type RepositoryI interface {
	// User logic
	GetUserByUsername(username string) (models.User, error)
	GetUserByUsernameAndPassword(username, password string) (models.User, error)
	CreateUser(u models.User) error

	// Articles logic
	CreateArticle(a models.Article) error
	GetAllArticles(userID int, title string) ([]models.Article, error)
	GetArticleByID(id int) (models.Article, error)
	UpdateArticle(a models.Article) error
	DeleteArticle(id int) error
}
