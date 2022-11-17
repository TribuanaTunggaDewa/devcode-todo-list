package dto

type ActivityGroupCreateDto struct {
	Email string `json:"email" form:"email" validate:"required"`
	Title string `json:"title" form:"title" validate:"required"`
}
