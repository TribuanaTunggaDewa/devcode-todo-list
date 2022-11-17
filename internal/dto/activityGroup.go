package dto

type ActivityGroupCreateDto struct {
	Email string `json:"email" validate:"required"`
	Title string `json:"title" validate:"required"`
}
