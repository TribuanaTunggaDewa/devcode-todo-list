package repository

import (
	"errors"
	"reflect"
	"todo-list/internal/abstractions"

	"gorm.io/gorm"
)

type TodoItemRepository struct {
	DBConnection *gorm.DB
	Model        any
}

func NewTodoItem(DBConnection *gorm.DB, Model any) *TodoItemRepository {
	return &TodoItemRepository{
		DBConnection,
		Model,
	}
}

func (r *TodoItemRepository) Find(payload *abstractions.GetQueries, out any, param string) error {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		return errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model)

	if param != "" {
		query = query.Where("activity_group_id = ?", param)
	}

	err := query.Find(out).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoItemRepository) FindById(id int, payload *abstractions.GetByIdQueries, out any) error {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		return errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model)

	err := query.Where("id = ?", id).First(out).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoItemRepository) Create(input any) error {
	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Ptr {
		return errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model).Create(input)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoItemRepository) Update(id int, payload any) error {
	t := reflect.TypeOf(payload)
	if t.Kind() != reflect.Ptr {
		return errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model)
	err := query.Where("id = ?", id).Updates(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoItemRepository) Delete(data any) error {
	query := r.DBConnection.Model(data)

	err := query.Delete(data).Error
	if err != nil {
		return err
	}

	return nil
}
