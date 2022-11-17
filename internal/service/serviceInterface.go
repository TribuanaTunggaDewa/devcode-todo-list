package service

import "todo-list/internal/repository"

type ServiceInteface interface {
	repository.RepositoryInterface
}
