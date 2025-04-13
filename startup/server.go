package startup

import (
	"app/arch/network"
	"app/arch/postgres"
	"app/config"
	_ "app/docs"
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

type Shutdown = func()

// @title Task Management API
// @version 1.0
// @description Task API documentation
// @host localhost:8080
// @BasePath /
func Start() {
	env := config.IniitEnv(".env", true)
	log.WithField("environment", env).Info("Environment loaded")
	shutdown, router, _ := create(env)
	defer shutdown()
	router.Start(env.ServerHost, env.ServerPort)
}

func create(env *config.Env) (Shutdown, network.Router, Module) {
	context := context.Background()

	//Connect database
	dbConfig := postgres.PostgresConfig{
		Host:        env.DBHost,
		Port:        env.DBPort,
		User:        env.DBUser,
		Pwd:         env.DBPassword,
		Name:        env.DBName,
		MinPoolSize: env.DBMinPoolSize,
		MaxPoolSize: env.DBMaxPoolSize,
		Timeout:     time.Duration(env.DBQueryTimeout) * time.Second,
	}
	db := postgres.CreateDatabase(context, dbConfig)
	db.Connect()

	// init module
	module := CreateModule(context, env, db)

	// init routes
	router := network.CreateNewRouter(env.GoMode)
	router.InitControllers(module.Controllers())

	shutdown := func() {
		db.Disconnect()
	}

	log.Info("Server started successfully")
	return shutdown, router, module
}
