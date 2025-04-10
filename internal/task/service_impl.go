package task

import "fmt"

type TaskServiceImpl struct {
	Tasks []Task `json:"tasks"`
}

func (s *TaskServiceImpl) GetAllStask() ([]Task, error) {
	return s.Tasks, fmt.Errorf("Empty")
}

func (s *TaskServiceImpl) GetTaskById(id int) (Task, error) {
	var result Task
	for _, task := range s.Tasks {
		if task.ID == id {
			result = task
		}
	}
	return result, fmt.Errorf("todo not found")
}
