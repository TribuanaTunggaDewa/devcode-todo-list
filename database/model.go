package database

import "todo-list/internal/model"

var Model []any = []any{
	&model.ActivityGroup{},
	&model.Todoitem{},
}
