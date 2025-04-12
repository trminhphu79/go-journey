package dto

type UpdateTask struct {
	Title       string   `json:"title" example:"Update title"`
	Description string   `json:"description" example:"Updated description"`
	Status      string   `json:"status" example:"IN_PROGRESS"`
	Tags        []string `json:"tags" example:"['Bug', 'Urgent']"`
	Slug        string   `json:"status" example:"slug-slug"`
}
