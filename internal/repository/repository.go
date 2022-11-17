package repository

import (
	"errors"
	"reflect"
	"todo-list/internal/abstractions"

	"gorm.io/gorm"
)

type Repository struct {
	DBConnection *gorm.DB
	Model        any
}

func NewRepository(DBConnection *gorm.DB, Model any) *Repository {
	return &Repository{
		DBConnection,
		Model,
	}
}

func (r *Repository) Find(payload *abstractions.GetQueries, out any) error {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		return errors.New("out must be a pointer")
	}

	query := r.DBConnection.Model(r.Model)

	err := query.Find(out).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindById(id int, payload *abstractions.GetByIdQueries, out any) error {
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

func (r *Repository) Create(input any) error {
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

func (r *Repository) Update(id int, payload any) error {
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

func (r *Repository) Delete(data any) error {
	query := r.DBConnection.Model(data)

	err := query.Delete(data).Error
	if err != nil {
		return err
	}

	return nil
}
