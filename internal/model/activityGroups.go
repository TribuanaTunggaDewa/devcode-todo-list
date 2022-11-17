package model

import (
	"time"
	"todo-list/internal/abstractions"

	"gorm.io/gorm"
)

type ActivityGroup struct {
	abstractions.Model
	Email string `json:"email"`
	Title string `json:"title"`
}

func (m *ActivityGroup) BeforeCreate(tx *gorm.DB) {
	m.CreatedAt = time.Now().In(time.Local)
}

func (m *ActivityGroup) BeforeUpdate(tx *gorm.DB) {
	m.UpdatedAt = time.Now().In(time.Local)
}
