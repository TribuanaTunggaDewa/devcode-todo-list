package model

import (
	"time"
	"todo-list/internal/abstractions"

	"gorm.io/gorm"
)

type Todo struct {
	abstractions.Model
	ActivityGroupId int    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        string `json:"is_active"`
	Priority        string `json:"priority"`
}

func (m *Todo) BeforeCreate(tx *gorm.DB) {
	m.CreatedAt = time.Now().In(time.Local)
}

func (m *Todo) BeforeUpdate(tx *gorm.DB) {
	m.UpdatedAt = time.Now().In(time.Local)
}
