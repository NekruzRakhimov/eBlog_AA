package service

import (
	"database/sql"
	"eBlog/internal/models"
	"errors"
)

func (s *Service) CreateArticle(a models.Article) error {
	return s.repository.CreateArticle(a)
}

func (s *Service) GetAllArticles(userID int, title string) ([]models.Article, error) {
	return s.repository.GetAllArticles(userID, title)
}

func (s *Service) GetArticleByID(id int) (models.Article, error) {
	article, err := s.repository.GetArticleByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Article{}, errors.New("статья c таким ID не найдена")
		}
		return models.Article{}, err
	}
	return article, nil
}

func (s *Service) UpdateArticle(a models.Article) error {
	_, err := s.repository.GetArticleByID(a.ID)
	if err != nil {
		return errors.New("статья c таким ID не найдена")
	}

	return s.repository.UpdateArticle(a)
}

func (s *Service) DeleteArticle(id int) error {
	err := s.repository.DeleteArticle(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("статья c таким ID не найдена")
		}
		return err
	}
	return nil
}
