package service

import (
	"todo-list/internal/abstractions"
	"todo-list/internal/repository"
)

type Service struct {
	reposiory repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) *Service {
	return &Service{
		reposiory: repository,
	}
}

func (s *Service) Find(payload *abstractions.GetQueries, out any, param string) (any, error) {

	result, err := s.reposiory.Find(payload, out, param)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) FindById(id int, payload *abstractions.GetByIdQueries, out any) (any, error) {

	result, err := s.reposiory.FindById(id, payload, out)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Create(payload any) (any, error) {

	result, err := s.reposiory.Create(payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *Service) Update(id int, payload any) (any, error) {

	result, err := s.reposiory.Update(id, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *Service) Delete(data any) (any, error) {

	result, err := s.reposiory.Delete(data)

	if err != nil {
		return nil, err
	}

	return result, nil

}
