package repository

import (
	"errors"
	"todo-list/internal/abstractions"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	Mock mock.Mock
}

func NewRepositoryMock(mock mock.Mock) *repositoryMock {
	return &repositoryMock{
		Mock: mock,
	}
}

func (r *repositoryMock) Find(payload *abstractions.GetQueries, out any, param string) (any, error) {
	arguments := r.Mock.Called(param)
	data := arguments.Get(0)
	if arguments.Get(0) == nil {
		return nil, errors.New("not found")
	} else {
		return data, nil
	}

}
func (r *repositoryMock) FindById(id int, payload *abstractions.GetByIdQueries, out any) (any, error) {
	arguments := r.Mock.Called(id)
	data := arguments.Get(0)
	if arguments.Get(0) == nil {
		return nil, nil
	} else {
		return data, nil
	}
}
func (r *repositoryMock) Create(payload any) (any, error) {
	arguments := r.Mock.Called(payload)
	data := arguments.Get(0)
	return data, nil
}
func (r *repositoryMock) Update(id int, payload any) (any, error) {
	arguments := r.Mock.Called(id, payload)
	data := arguments.Get(0)
	return data, nil
}
func (r *repositoryMock) Delete(data any) (any, error) {
	arguments := r.Mock.Called(data)
	payload := arguments.Get(0)
	return payload, nil
}
