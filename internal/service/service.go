package service

import (
	"eBlog/internal/contracts"
	"eBlog/internal/repository"
)

type Service struct {
	repository contracts.RepositoryI
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}
