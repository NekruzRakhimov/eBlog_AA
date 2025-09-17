package service

import (
	"eBlog/internal/errs"
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
		if errors.Is(err, errs.ErrNotFound) {
			return models.Article{}, errs.ErrArticleNotFound
		}
		return models.Article{}, err
	}
	return article, nil
}

func (s *Service) UpdateArticle(a models.Article) error {
	_, err := s.repository.GetArticleByID(a.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrArticleNotFound
		}
		return err
	}

	return s.repository.UpdateArticle(a)
}

func (s *Service) DeleteArticle(id int) error {
	err := s.repository.DeleteArticle(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrArticleNotFound
		}
		return err
	}

	return nil
}
