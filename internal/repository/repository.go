package repository

import (
	"errors"
	"reflect"
	"todo-list/internal/abstractions"

	"gorm.io/gorm"
)

type repository struct {
	DBConnection *gorm.DB
	Model        any
}

func NewRepository(DBConnection *gorm.DB, Model any) *repository {
	return &repository{
		DBConnection,
		Model,
	}
}

func (r *repository) Find(payload *abstractions.GetQueries, out any, param string) (any, error) {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model)

	err := query.Find(out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *repository) FindById(id int, payload *abstractions.GetByIdQueries, out any) (any, error) {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model)

	err := query.Where("id = ?", id).First(out).Error
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *repository) Create(input any) (any, error) {
	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model).Create(input)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *repository) Update(id int, payload any) (any, error) {
	t := reflect.TypeOf(payload)
	if t.Kind() != reflect.Ptr {
		return nil, errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model)
	err := query.Where("id = ?", id).Updates(payload).Error
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (r *repository) Delete(data any) (any, error) {
	query := r.DBConnection.Model(data)

	err := query.Delete(data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
