package service

import (
	"testing"
	"todo-list/internal/abstractions"
	"todo-list/internal/model"
	"todo-list/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var reposioryMock = repository.NewRepositoryMock(mock.Mock{})
var serviceTest = NewService(reposioryMock)

func TestService_Find_Success(t *testing.T) {

	datas := []model.Activitie{
		{
			Model: abstractions.Model{
				ID: 1,
			},
			Email: "nugroho@gmail.com",
			Title: "Yugroh",
		},
		{
			Model: abstractions.Model{
				ID: 2,
			},
			Email: "yusril@gmail.com",
			Title: "busroh",
		},
		{
			Model: abstractions.Model{
				ID: 1,
			},
			Email: "ela@gmail.com",
			Title: "subroh",
		},
	}

	test := []struct {
		name    string
		data    []model.Activitie
		param   string
		payload abstractions.GetQueries
	}{{
		name:    "test-1",
		param:   "",
		data:    datas,
		payload: abstractions.GetQueries{},
	}, {
		name:    "test-2",
		param:   "",
		data:    datas,
		payload: abstractions.GetQueries{},
	},
		{
			name:    "test-3",
			param:   "",
			data:    datas,
			payload: abstractions.GetQueries{},
		},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			reposioryMock.Mock.On("Find", v.param).Return(v.data)
			result, err := serviceTest.Find(&abstractions.GetQueries{}, &model.Activitie{}, "")
			conv := result.([]model.Activitie)
			assert.Nil(t, err)
			assert.NotNil(t, conv)
		})
	}

}
