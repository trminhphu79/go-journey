package repository

import (
	"app/api/task/dto"
	"app/api/task/model"
	"app/arch/network"
	"app/arch/postgres"
	"errors"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ITaskRepo interface {
	FindOne(id string) (*model.Task, error)
	FindOneAndUpdate(id uuid.UUID, dto dto.UpdateTask) (*model.Task, error)
	Delete(id string) (string, error)
}

type taskRepo struct {
	db postgres.Database
}

func CreateTaskRepository(db postgres.Database) ITaskRepo {
	return &taskRepo{}
}

func (r *taskRepo) FindOne(id string) (*model.Task, error) {
	return nil, nil
}

func (r *taskRepo) FindOneAndUpdate(id uuid.UUID, dto dto.UpdateTask) (*model.Task, error) {
	task := model.Task{
		ID:          id,
		Title:       dto.Title,
		Slug:        dto.Slug,
		Description: dto.Description,
	}
	result := r.db.GetInstance().Save(task)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logger.Error("Task not found ", result.Error)
		return nil, result.Error
	}

	if result.Error != nil {
		return nil, network.NewInternalServerErr("Update task failed", result.Error)
	}

	if result.RowsAffected <= 0 {
		logger.WithField("taskId", id).Warn("Failed to update task, not found")
		return nil, network.NewNotFoundErr("Update task failed", nil)
	}

	return &task, nil
}

func (r *taskRepo) Delete(id string) (string, error) {
	return "nil", nil
}
