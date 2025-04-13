package task

import (
	"app/api/task/dto"
	"app/arch/network"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
	group.DELETE("/:id", c.deleteTask)
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
	log.Info("Controller getTaskById")
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
	log.Info("Controller createNewTask")
	var input dto.CreateTaskDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.WithError(err).Warn("Invalid input for task update")
		c.Send(ctx).BadRequestErr("Input value is invalid", err)
		return
	}
	value, err := c.service.CreateTask(input)
	if err != nil {
		c.Send(ctx).MixedError(err)
		return
	}

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
	idStr := ctx.Param("id")
	log.WithField("taskId", idStr).Info("Processing task update request")

	// (optional) convert to UUID if needed
	taskID, err := uuid.Parse(idStr)
	if err != nil {
		log.WithError(err).Warn("Invalid UUID format for task ID")
		c.Send(ctx).BadRequestErr("Invalid UUID format for task ID", err)
		return
	}

	var input dto.UpdateTask
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.WithError(err).Warn("Invalid input for task update")
		c.Send(ctx).BadRequestErr("Input value is invalid", err)
		return
	}

	task, err := c.service.UpdateTask(taskID, input)

	if err != nil {
		log.WithError(err).Error("Failed to update task")
		c.Send(ctx).MixedError(err)
		return
	}

	log.WithField("taskId", idStr).Info("Task updated successfully")
	c.Send(ctx).SuccessDataRes("update success", task)
}

// deleteTask godoc
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID (UUID)"
// @Success 200 {object} network.response{data=int64} "Success response with number of affected rows"
// @Router /api/v1/task/{id} [delete]
func (c *taskController) deleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	taskID, err := uuid.Parse(idStr)
	if err != nil {
		log.WithError(err).Warn("Invalid UUID format for task ID")
		c.Send(ctx).BadRequestErr("Invalid UUID format for task ID", err)
		return
	}
	affected, err := c.service.DeleteTask(taskID)
	if err != nil {
		log.WithError(err).Error("Failed to delete task")
		c.Send(ctx).MixedError(err)
		return
	}
	log.WithFields(log.Fields{
		"taskId":   idStr,
		"affected": affected,
	}).Info("Task deleted successfully")
	c.Send(ctx).SuccessDataRes("Delete task success", affected)
}

func (c *taskController) pagingTask(ctx *gin.Context) {
	log.WithFields(
		log.Fields{
			"keyword": ctx.Param("keyword"),
			"offset": ctx.Param("offset"),
			"limit": ctx.Param("limit"),
		}
	).Info("Paging Task")
}