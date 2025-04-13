package dto

type CreateTaskDTO struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Status      string   `json:"status"`
	Slug        string   `json:"slug"`
}
