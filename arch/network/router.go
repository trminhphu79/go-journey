package network

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type router struct {
	engine *gin.Engine
}

func (r *router) GetEngine() *gin.Engine {
	return r.engine
}

func (r *router) Start(ip string, port uint16) {
	address := fmt.Sprintf("%s:%d", ip, port)
	r.engine.Run(address)
}

func CreateNewRouter(mode string) Router {
	gin.SetMode(mode)
	eng := gin.Default()
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r := router{
		engine: eng,
	}
	return &r
}

func (r *router) InitControllers(controllers []Controller) {
	for _, c := range controllers {
		g := r.engine.Group(c.Path())
		c.AddRouters(g)
	}
}
