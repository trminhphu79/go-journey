package task

import (
	"app/api/task/dto"
	"app/arch/network"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	group.PATCH("/:id", c.updateTask)
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
	task := c.service.FindTaskById("1123")
	c.Send(ctx).SuccessDataRes("Get Success Full", task)
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
	value, _ := c.service.CreateTask()
	c.Send(ctx).SuccessDataRes("Create Success Full", value)
}

// updateTask godoc
// @Summary Update a task
// @Description Update task fields partially by ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID (UUID)"
// @Param input body dto.UpdateTask true "Updated task fields"
// @Success 200 {object} model.Task
// @Router /api/v1/task/{id} [patch]
func (c *taskController) updateTask(ctx *gin.Context) {
	fmt.Println("ctx.Param: ", ctx.Request.URL)
	idStr := ctx.Param("id")

	fmt.Println("Task ID:", idStr)

	// (optional) convert to UUID if needed
	taskID, err := uuid.Parse(idStr)
	if err != nil {
		c.Send(ctx).BadRequestErr("Invalid UUID format for task ID", err)
		return
	}
	var input dto.UpdateTask
	if err := ctx.ShouldBindJSON(&input); err != nil {
		c.Send(ctx).BadRequestErr("Input value is invalid", err)
		return
	}

	task, err := c.service.UpdateTask(taskID, input)

	if err != nil {
		c.Send(ctx).BadRequestErr("Internal server error", err)
		return
	}

	c.Send(ctx).SuccessDataRes("update success", task)
}
