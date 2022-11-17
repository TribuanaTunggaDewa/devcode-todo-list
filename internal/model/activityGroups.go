package model

import "todo-list/internal/abstractions"

type ActivityGroup struct {
	abstractions.Model
	Email string `json:"email"`
	Title string `json:"title"`
}
