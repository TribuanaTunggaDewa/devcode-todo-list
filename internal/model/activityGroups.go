package model

import (
	"time"
	"todo-list/internal/abstractions"

	"gorm.io/gorm"
)

type Activitie struct {
	abstractions.Model
	Email string `json:"email"`
	Title string `json:"title"`
}

func (m *Activitie) BeforeCreate(tx *gorm.DB) {
	m.CreatedAt = time.Now().In(time.Local)
}

func (m *Activitie) BeforeUpdate(tx *gorm.DB) {
	m.UpdatedAt = time.Now().In(time.Local)
}
