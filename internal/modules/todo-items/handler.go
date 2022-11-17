package todoitems

import (
	"todo-list/database"
	"todo-list/internal/abstractions"
	"todo-list/internal/model"
	"todo-list/internal/repository"
	"todo-list/internal/service"
	"todo-list/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service.Service
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Get(c echo.Context) error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return err
	}

	data := new([]model.Todoitem)
	repository := repository.NewRepository(dbconnection, data)
	h.service = service.NewService(repository)

	err = h.service.Repository.Find(&abstractions.GetQueries{}, data)

	if err != nil {
		return err
	}

	return response.CustomSuccessBuilder(200, data, "success", "success").Send(c)
}

func (h *handler) Store(c echo.Context) error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return err
	}

	todoItem := new(model.Todoitem)

	repository := repository.NewRepository(dbconnection, todoItem)
	h.service = service.NewService(repository)

	err = h.service.Repository.Create(todoItem)

	if err != nil {
		return err
	}

	return response.CustomSuccessBuilder(200, todoItem, "success", "success").Send(c)

}

func (h *handler) Update(c echo.Context) error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return err
	}

	todoItem := new(model.Todoitem)

	repository := repository.NewRepository(dbconnection, todoItem)
	h.service = service.NewService(repository)

	err = h.service.Repository.Update(3, todoItem)

	if err != nil {
		return err
	}

	return response.CustomSuccessBuilder(200, todoItem, "success", "success").Send(c)
}

func (h *handler) Delete(c echo.Context) error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return err
	}

	todoItem := &model.Todoitem{
		Model: abstractions.Model{
			ID: 3,
		},
	}

	repository := repository.NewRepository(dbconnection, todoItem)
	h.service = service.NewService(repository)

	err = h.service.Repository.Delete(todoItem)
	if err != nil {
		return err
	}

	return response.CustomSuccessBuilder(200, todoItem, "success", "success").Send(c)
}
