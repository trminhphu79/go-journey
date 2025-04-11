package task

import (
	"app/api/task/model"
	"app/arch/network"
	"app/arch/postgres"
	"fmt"
)

type TaskService interface {
	FindTaskById(taskId string) *model.Task
	CreateTask(task *model.Task) (*model.TaskStatus, error)
}
type taskService struct {
	network.BaseService
	db postgres.Database
}

func CreateService(db postgres.Database) TaskService {
	return &taskService{
		BaseService: network.NewBaseService(),
		db:          db.GetInstance(),
	}
}

func (s *taskService) FindTaskById(taskId string) *model.Task {
	fmt.Println("Find task by id: ", taskId)
	return nil
}

func (s *taskService) CreateTask(task *model.Task) (*model.TaskStatus, error) {
	fmt.Println("Create task: ", task)
	status := model.Todo
	return &status, nil
}
