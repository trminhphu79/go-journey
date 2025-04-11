package task

import (
	"app/arch/network"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	network.BaseController
}

func (c *taskController) AddRouters(group *gin.RouterGroup) {
	group.GET("/:id", c.getTaskById)
	group.POST("/", c.createNewTask)
}

func (c *taskController) getTaskById(ctx *gin.Context) {

}

func (c *taskController) createNewTask(ctx *gin.Context) {}
