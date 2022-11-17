package repository

import "todo-list/internal/abstractions"

type RepositoryInterface interface {
	Find(payload *abstractions.GetQueries, out any) error
	FindById(payload *abstractions.GetByIdQueries, out any) error
	Create(payload any) error
	Update(id int, payload any) error
	Delete(data any) error
}
