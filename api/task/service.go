package task

import (
	"app/api/task/dto"
	"app/api/task/model"
	"app/arch/network"
	"app/arch/postgres"

	"errors"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TaskService interface {
	FindTaskById(taskId string) *model.Task
	CreateTask(dto.CreateTaskDTO) (*model.Task, error)
	UpdateTask(taskId uuid.UUID, input dto.UpdateTask) (task *model.Task, err error)
	DeleteTask(taskId uuid.UUID) (affected int64, err error)
	PagingTask(dto.PagingTaskDto) ([]model.Task, error)
	AssignTask(dto.AssignTaskDto) (msg string, err error)
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
	log.WithField("taskId", taskId).Info("Finding task by ID")

	return nil
}

func (s *taskService) CreateTask(input dto.CreateTaskDTO) (*model.Task, error) {
	task := model.Task{
		Title:        input.Title,
		Description:  input.Description,
		Status:       model.Todo,
		Slug:         input.Slug,
		AssigneeID:   nil,
		AssignedByID: nil,
	}

	result := s.db.GetInstance().Create(&task)
	if result.Error != nil {
		log.WithError(result.Error).Error("Failed to create task")
		return nil, result.Error
	}

	log.WithField("rowsAffected", result.RowsAffected).Info("Task created successfully")
	return &task, nil
}

func (s *taskService) UpdateTask(taskId uuid.UUID, input dto.UpdateTask) (returnValue *model.Task, err error) {
	log.WithFields(log.Fields{
		"taskId": taskId,
		"input":  input,
	}).Info("Updating task")

	task := model.Task{
		ID:          taskId,
		Title:       input.Title,
		Slug:        input.Slug,
		Description: input.Description,
	}

	result := s.db.GetInstance().Save(task)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("Task not found ", result.Error)
		return nil, result.Error
	}

	if result.Error != nil {
		return nil, network.NewInternalServerErr("Update task failed", result.Error)
	}

	if result.RowsAffected <= 0 {
		log.WithField("taskId", taskId).Warn("Failed to update task, not found")
		return nil, network.NewNotFoundErr("Update task failed", err)
	}

	log.WithField("taskId", taskId).Info("Task updated successfully")
	return &task, nil

}

func (s *taskService) DeleteTask(taskId uuid.UUID) (affected int64, err error) {
	result := s.db.GetInstance().Where("id = ?", taskId).Delete(&model.Task{})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("Task not found ", result.Error)
		return 0, result.Error
	}

	if result.Error != nil {
		return 0, network.NewInternalServerErr("Delete task failed", result.Error)
	}

	if result.RowsAffected <= 0 {
		return result.RowsAffected, network.NewNotFoundErr("Task not found", err)
	}
	return result.RowsAffected, nil
}

// Implement the method in your service
func (s *taskService) PagingTask(input dto.PagingTaskDto) ([]model.Task, error) {
	log.WithFields(log.Fields{
		"keyword": input.Keyword,
		"offset":  input.Offset,
		"limit":   input.Limit,
		"status":  input.Status,
	}).Info("Paging tasks from database")

	var tasks []model.Task
	query := s.db.GetInstance().Model(&model.Task{})

	if input.Keyword != "" {
		query = query.Where("title LIKE ?", "%"+input.Keyword+"%")
	}

	if input.Status != "" {
		query = query.Where("status = ?", input.Status)
	}

	result := query.Limit(int(input.Limit)).Offset(int(input.Offset)).Find(&tasks)
	if result.Error != nil {
		log.Error("Failed to retrieve paging task ", result.Error)
		return nil, result.Error
	}

	return tasks, nil
}

func (s *taskService) AssignTask(input dto.AssignTaskDto) (msg string, err error) {
	return "", nil
}
