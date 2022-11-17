package model

import "todo-list/internal/abstractions"

type Todoitem struct {
	abstractions.Model
	ActivityGroupId string `json:"activity_group_id"`
	Title           string `json:"title"`
}
