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
	authProvider network.AuthenticationProvider,
) network.Controller {
	return &taskController{
		BaseController: network.NewBaseController("api/v1/task", authProvider),
		service:        service,
	}
}

func (c *taskController) AddRouters(group *gin.RouterGroup) {
	group.GET("/:id", c.getTaskById)
	group.POST("/", c.createNewTask)
	group.PATCH("/:id", c.updateTask)
	group.DELETE("/:id", c.deleteTask)
	group.POST("/paging", c.pagingTask)
	group.POST("/assign", c.assignTask)
}

// GetTaskById godoc
// @Summary Get a task by ID
// @Description Get a task by its ID
// @Tags Task
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
// @Tags Task
// @Accept json
// @Produce json
// @Body task body model.Task true "Task object"
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
		c.Send(ctx).ComposeError(err)
		return
	}

	c.Send(ctx).SuccessDataRes("Create Success Full", value)
}

// updateTask godoc
// @Summary Update a task
// @Description Update task fields partially by ID
// @Tags Task
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
		c.Send(ctx).ComposeError(err)
		return
	}

	log.WithField("taskId", idStr).Info("Task updated successfully")
	c.Send(ctx).SuccessDataRes("update success", task)
}

// deleteTask godoc
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags Task
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
		c.Send(ctx).ComposeError(err)
		return
	}
	log.WithFields(log.Fields{
		"taskId":   idStr,
		"affected": affected,
	}).Info("Task deleted successfully")
	c.Send(ctx).SuccessDataRes("Delete task success", affected)
}

// PagingTask godoc
// @Summary Get paginated list of tasks
// @Description Retrieves a paginated list of tasks filtered by title and status
// @Tags Task
// @Accept json
// @Produce json
// @Param request body dto.PagingTaskDto true "Pagination and filter criteria"
// @Success 200 {object} network.response{data=[]model.Task} "Success response with paginated tasks"
// @Failure 400 {object} network.response "Bad request error"
// @Failure 500 {object} network.response "Internal server error"
// @Router /api/v1/task/paging [post]
func (c *taskController) pagingTask(ctx *gin.Context) {
	var input dto.PagingTaskDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.WithError(err).Warn("Invalid input for task paging")
		c.Send(ctx).BadRequestErr("Input value is invalid", err)
		return
	}

	log.WithFields(log.Fields{
		"keyword": input.Keyword,
		"offset":  input.Offset,
		"limit":   input.Limit,
		"status":  input.Status,
	}).Info("Processing task paging request")

	// You need to implement this in your service
	tasks, err := c.service.PagingTask(input)
	if err != nil {
		log.WithError(err).Error("Failed to retrieve paged tasks")
		c.Send(ctx).BadRequestErr("Failed to retrieve tasks", err)
		return
	}

	c.Send(ctx).SuccessDataRes("Tasks retrieved successfully", tasks)
}

// Assign Task godoc
// @Summary Assign task to a user
// @Description Assign task to a user
// @Tags Task
// @Accept json
// @Produce json
// @Param request body dto.PagingTaskDto true "Pagination and filter criteria"
// @Success 200 {object} network.response{data=model.Task} "Success response with paginated tasks"
// @Failure 400 {object} network.response "Bad request error"
// @Failure 500 {object} network.response "Internal server error"
// @Router /api/v1/task/assign [post]
func (c *taskController) assignTask(ctx *gin.Context) {
	c.Send(ctx).SuccessMsgRes("Assign success")
}
