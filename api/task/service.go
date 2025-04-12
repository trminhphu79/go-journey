package task

import (
	"app/api/task/dto"
	"app/api/task/model"
	"app/arch/network"
	"app/arch/postgres"
	"fmt"

	"github.com/google/uuid"
)

type TaskService interface {
	FindTaskById(taskId string) *model.Task
	CreateTask() (*model.Task, error)
	UpdateTask(taskId uuid.UUID, input dto.UpdateTask) (task *model.Task, err error)
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

func (s *taskService) CreateTask() (*model.Task, error) {
	task := model.Task{
		Title:       "task A",
		Description: "Description",
		Status:      model.Todo,
		Slug:        "create-new-task",
	}

	result := s.db.GetInstance().Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("result.RowsAffected: ", result.RowsAffected)
	return &task, nil
}

func (s *taskService) UpdateTask(taskId uuid.UUID, input dto.UpdateTask) (returnValue *model.Task, err error) {
	task := model.Task{
		ID:          taskId,
		Title:       input.Title,
		Slug:        input.Slug,
		Description: input.Description,
	}

	result := s.db.GetInstance().Save(task)

	if result.RowsAffected > 0 {
		return &task, nil
	}

	return nil, network.NewNotFoundErr("Update task failed", err)
}
