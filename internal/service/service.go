package service

import "todo-list/internal/repository"

type Service struct {
	repository *repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) *Service {
	return &Service{
		repository: &repository,
	}
}
