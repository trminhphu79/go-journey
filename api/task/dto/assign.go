package dto

type AssignTaskDto struct {
	ID           string `json:"id"  binding:"required"`
	AssigneeId   string `json:"assigneeId"  binding:"required"`
	AssignedById string `json:"assignedId"  binding:"required"`
}
