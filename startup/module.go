package startup

import (
	"app/arch/network"
	"app/config"
	"context"
)

type Module network.Module[module]

type module struct {
	Context context.Context
	Env     *config.Env
}

func (m *module) GetInstance() *module {
	return m
}

func (m *module) Controllers() []network.Controller {
	return []network.Controller{}
}
