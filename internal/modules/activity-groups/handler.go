package activitygroups

import (
	"todo-list/database"
	"todo-list/internal/abstractions"
	"todo-list/internal/model"
	"todo-list/internal/repository"
	"todo-list/internal/service"
)

type handler struct {
	service *service.Service
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Get() error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return err
	}

	data := new([]*model.ActivityGroup)
	repository := repository.NewRepository(dbconnection, data)
	h.service = service.NewService(repository)

	err = h.service.Repository.Find(&abstractions.GetQueries{}, data)

	if err != nil {
		return err
	}

	return nil
}

func (h *handler) Store() error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return err
	}

	activityGroup := &model.ActivityGroup{
		Email: "tes@gmail.com",
		Title: "test insert",
	}

	repository := repository.NewRepository(dbconnection, activityGroup)
	h.service = service.NewService(repository)

	err = h.service.Repository.Create(activityGroup)

	if err != nil {
		return err
	}

	return nil

}

func (h *handler) Update() error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return err
	}

	activityGroup := &model.ActivityGroup{
		Email: "tes@gmail.com",
		Title: "test update",
	}

	repository := repository.NewRepository(dbconnection, activityGroup)
	h.service = service.NewService(repository)

	err = h.service.Repository.Update(3, activityGroup)

	if err != nil {
		return err
	}

	return nil
}

func (h *handler) Delete() error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return err
	}

	activityGroup := &model.ActivityGroup{
		Model: abstractions.Model{
			ID: 3,
		},
	}

	repository := repository.NewRepository(dbconnection, activityGroup)
	h.service = service.NewService(repository)

	err = h.service.Repository.Delete(activityGroup)
	if err != nil {
		return err
	}

	return nil
}
