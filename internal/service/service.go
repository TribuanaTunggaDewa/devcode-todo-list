package service

import "todo-list/internal/repository"

type Service struct {
	Repository repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) *Service {
	return &Service{
		Repository: repository,
	}
}
