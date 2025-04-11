package startup

import (
	"app/arch/network"
	"app/config"
	"fmt"
)

type Shutdown = func()

func Start() {
	env := config.IniitEnv(".env", true)
	fmt.Println("ENV: ", env)
	shutdown, router := create(env)
	defer shutdown()
	router.Start(env.ServerHost, env.ServerPort)
}

func create(env *config.Env) (Shutdown, network.Router) {
	shutdown := func() {

	}

	fmt.Println("Start server succesfully!!")
	router := network.CreateNewRouter(env.GoMode)
	return shutdown, router
}
