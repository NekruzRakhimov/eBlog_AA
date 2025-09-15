package service

import (
	"database/sql"
	"eBlog/internal/models"
	"eBlog/internal/repository"
	"errors"
)

func CreateArticle(a models.Article) error {
	return repository.CreateArticle(a)
}

func GetAllArticles(userID int, title string) ([]models.Article, error) {
	return repository.GetAllArticles(userID, title)
}

func GetArticleByID(id int) (models.Article, error) {
	article, err := repository.GetArticleByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Article{}, errors.New("статья c таким ID не найдена")
		}
		return models.Article{}, err
	}
	return article, nil
}

func UpdateArticle(a models.Article) error {
	_, err := repository.GetArticleByID(a.ID)
	if err != nil {
		return errors.New("статья c таким ID не найдена")
	}

	return repository.UpdateArticle(a)
}

func DeleteArticle(id int) error {
	err := repository.DeleteArticle(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("статья c таким ID не найдена")
		}
		return err
	}
	return nil
}
