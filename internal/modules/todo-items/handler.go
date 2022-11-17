package todoitems

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"todo-list/database"
	"todo-list/internal/abstractions"
	"todo-list/internal/dto"
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

	param := c.QueryParam("activity_group_id")

	if err != nil {
		return response.CustomErrorBuilder(500, &dto.ErrorNilObject{}, "error", "error").Send(c)
	}
	data := new([]model.Todoitem)
	repository := repository.NewTodoItem(dbconnection, data)
	h.service = service.NewService(repository)

	err = h.service.Repository.Find(&abstractions.GetQueries{}, data, param)

	if err != nil {
		return response.ErrorBuilder(response.Constant.Error.NotFound, errors.New("activity-group not found")).Send(c)
	}

	return response.CustomSuccessBuilder(200, data, "success", "success").Send(c)
}

func (h *handler) GetById(c echo.Context) error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return response.CustomErrorBuilder(500, &dto.ErrorNilObject{}, "error", "error").Send(c)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorBuilder(response.Constant.Error.BadRequest, err).Send(c)
	}

	data := new([]model.Todoitem)
	repository := repository.NewTodoItem(dbconnection, data)
	h.service = service.NewService(repository)

	err = h.service.Repository.FindById(id, &abstractions.GetByIdQueries{}, data)

	if err != nil {
		return response.CustomErrorBuilder(404, &dto.ErrorNilObject{}, fmt.Sprintf("Activity with ID %v Not Found", id), "Not Found").Send(c)

	}

	return response.CustomSuccessBuilder(200, data, "success", "success").Send(c)
}

func (h *handler) Store(c echo.Context) error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return response.CustomErrorBuilder(500, &dto.ErrorNilObject{}, "error", "error").Send(c)
	}

	todoItem := new(model.Todoitem)

	repository := repository.NewTodoItem(dbconnection, todoItem)
	h.service = service.NewService(repository)

	payload := new(dto.ItemTodoCreateDto)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(response.Constant.Error.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.CustomErrorBuilder(400, dto.ErrorNilObject{}, err.Error(), "Bad Request").Send(c)
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		return response.ErrorBuilder(response.Constant.Error.InternalServerError, err).Send(c)
	}

	err = json.Unmarshal(bytes, &todoItem)
	if err != nil {
		return response.ErrorBuilder(response.Constant.Error.InternalServerError, err).Send(c)
	}

	todoItem.IsActive = "1"
	todoItem.Priority = "very-high"

	err = h.service.Repository.Create(todoItem)

	if err != nil {
		return response.ErrorBuilder(response.Constant.Error.InternalServerError, err).Send(c)

	}

	return response.CustomSuccessBuilder(200, todoItem, "success", "success").Send(c)

}

func (h *handler) Update(c echo.Context) error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return response.CustomErrorBuilder(500, &dto.ErrorNilObject{}, "error", "error").Send(c)
	}

	todoItem := new(model.Todoitem)

	repository := repository.NewTodoItem(dbconnection, todoItem)
	h.service = service.NewService(repository)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorBuilder(response.Constant.Error.BadRequest, err).Send(c)
	}

	payload := new(dto.ItemTodoCreateDto)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(response.Constant.Error.BadRequest, err).Send(c)
	}

	if payload.ActivityGroupId == 0 && payload.Title == "" {
		return response.CustomErrorBuilder(400, dto.ErrorNilObject{}, "title cannot be null", "Bad Request").Send(c)
	}

	err = h.service.Repository.FindById(id, &abstractions.GetByIdQueries{}, todoItem)

	if err != nil {
		return response.CustomErrorBuilder(404, &dto.ErrorNilObject{}, fmt.Sprintf("Activity with ID %v Not Found", id), "Not Found").Send(c)

	}

	if payload.ActivityGroupId != 0 {
		todoItem.ActivityGroupId = payload.ActivityGroupId
	}

	if payload.Title != "" {
		todoItem.Title = payload.Title
	}

	err = h.service.Repository.Update(3, todoItem)

	if err != nil {
		return response.CustomErrorBuilder(500, dto.ErrorNilObject{}, err.Error(), response.Constant.Error.InternalServerError.Error()).Send(c)

	}

	return response.CustomSuccessBuilder(200, todoItem, "success", "success").Send(c)
}

func (h *handler) Delete(c echo.Context) error {
	dbconnection, err := database.GetConnection()

	if err != nil {
		return response.CustomErrorBuilder(500, &dto.ErrorNilObject{}, "error", "error").Send(c)
	}

	todoItem := new(model.Todoitem)

	repository := repository.NewTodoItem(dbconnection, todoItem)
	h.service = service.NewService(repository)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorBuilder(response.Constant.Error.BadRequest, err).Send(c)
	}

	err = h.service.Repository.FindById(id, &abstractions.GetByIdQueries{}, todoItem)

	if err != nil {
		return response.CustomErrorBuilder(404, &dto.ErrorNilObject{}, fmt.Sprintf("Activity with ID %v Not Found", id), "Not Found").Send(c)

	}

	err = h.service.Repository.Delete(todoItem)
	if err != nil {
		return err
	}

	return response.CustomSuccessBuilder(200, dto.ErrorNilObject{}, "success", "success").Send(c)

}
