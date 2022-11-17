package abstractions

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement;" param:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
