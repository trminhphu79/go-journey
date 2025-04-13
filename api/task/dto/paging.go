package dto

type PagingTaskDto struct {
	Keyword string `json:"keyword" example: "Js task"`
	Status  string `json:"status" example: "TODO"`
	Offset  int8   `json:"offset" example: 0`
	Limit   int8   `json:"limit" example: 5`
}
