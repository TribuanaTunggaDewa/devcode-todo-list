package dto

type ItemTodoCreateDto struct {
	ActivityGroupId int    `json:"activity_group_id" form:"activity_group_id" validate:"required"`
	Title           string `json:"title" form:"title" validate:"required"`
}
