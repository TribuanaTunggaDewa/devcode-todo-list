package service

import "todo-list/internal/abstractions"

type ServiceInteface interface {
	Find(payload *abstractions.GetQueries, out any, param string) (any, error)
	FindById(id int, payload *abstractions.GetByIdQueries, out any) (any, error)
	Create(payload any) (any, error)
	Update(id int, payload any) (any, error)
	Delete(data any) (any, error)
}
