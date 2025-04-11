package task

import (
	"app/arch/network"
	"fmt"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	network.BaseController
	service TaskService
}

func CreateController(
	service TaskService,
) network.Controller {
	return &taskController{
		BaseController: network.NewBaseController("api/v1/task"),
		service:        service,
	}
}

func (c *taskController) AddRouters(group *gin.RouterGroup) {
	group.GET("/:id", c.getTaskById)
	group.POST("/", c.createNewTask)
}

// GetTaskById godoc
// @Summary Get a task by ID
// @Description Get a task by its ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} model.Task
// @Failure 400 {object} network.apiError
// @Failure 404 {object} network.apiError
// @Router /api/v1/task/{id} [get]
func (c *taskController) getTaskById(ctx *gin.Context) {
	fmt.Println("Controller getTaskById")
	c.service.FindTaskById("1123")
	c.Send(ctx).SuccessMsgRes("Get Success Full")
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the input data
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body model.Task true "Task object"
// @Success 201 {object} model.TaskStatus
// @Failure 400 {object} network.apiError
// @Router /api/v1/task [post]
func (c *taskController) createNewTask(ctx *gin.Context) {
	fmt.Println("Controller createNewTask")
	c.Send(ctx).SuccessMsgRes("Create Success Full")
}
