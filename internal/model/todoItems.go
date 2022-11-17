package model

import (
	"time"
	"todo-list/internal/abstractions"

	"gorm.io/gorm"
)

type Todoitem struct {
	abstractions.Model
	ActivityGroupId string `json:"activity_group_id"`
	Title           string `json:"title"`
}

func (m *Todoitem) BeforeCreate(tx *gorm.DB) {
	m.CreatedAt = time.Now().In(time.Local)
}

func (m *Todoitem) BeforeUpdate(tx *gorm.DB) {
	m.UpdatedAt = time.Now().In(time.Local)
}
