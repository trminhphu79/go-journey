package framework

import (
	"app/api/auth"
	"app/api/auth/middleware"
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
	AuthService auth.AuthService
}

func (m *module) GetInstance() *module {
	return m
}

func (m *module) Controllers() []network.Controller {
	return []network.Controller{
		task.CreateController(m.TaskService, m.AuthenticationProvider()),
		auth.CreateController(m.AuthService, m.AuthenticationProvider()),
	}
}

func (m *module) AuthenticationProvider() network.AuthenticationProvider {
	return middleware.NewAuthenticateHandler(m.AuthService)
}
func CreateModule(context context.Context, env *config.Env, db postgres.Database) Module {
	taskService := task.CreateService(db)
	authService := auth.CreateAuthService(db, env)

	return &module{
		Context:     context,
		Env:         env,
		TaskService: taskService,
		AuthService: authService,
	}
}
