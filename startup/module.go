package startup

import (
	"app/api/task"
	"app/arch/network"
	"app/arch/postgres"
	"app/config"
	"context"
)

type Module network.Module[module]

type module struct {
	Context     context.Context
	Env         *config.Env
	TaskService task.TaskService
}

func (m *module) GetInstance() *module {
	return m
}

func (m *module) Controllers() []network.Controller {
	return []network.Controller{
		task.CreateController(m.TaskService),
	}
}

func CreateModule(context context.Context, env *config.Env, db postgres.Database) Module {
	taskService := task.CreateService(db)

	return &module{
		Context:     context,
		Env:         env,
		TaskService: taskService,
	}
}
