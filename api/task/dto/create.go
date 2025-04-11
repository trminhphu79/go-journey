package dto

type CreateTask struct {
	Title       string   `validate:"required"`
	Description string   `validate:"required"`
	Tags        []string `validate:"required"`
	Slug        string   `validate:"required"`
}
