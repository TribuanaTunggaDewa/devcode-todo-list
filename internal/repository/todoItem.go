package repository

import (
	"errors"
	"reflect"
	"todo-list/internal/abstractions"

	"gorm.io/gorm"
)

type TodoItemRepository struct {
	*repository
}

func NewTodoItem(DBConnection *gorm.DB, Model any) *TodoItemRepository {
	repo := NewRepository(DBConnection, Model)
	return &TodoItemRepository{
		repository: repo,
	}
}

func (r *TodoItemRepository) Find(payload *abstractions.GetQueries, out any, param string) (any, error) {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("out must be a pointer")
	}

	query := r.repository.DBConnection.Model(r.repository.Model)

	if param != "" {
		query = query.Where("activity_group_id = ?", param)
	}

	err := query.Find(out).Error
	if err != nil {
		return nil, err
	}

	return &out, nil
}
