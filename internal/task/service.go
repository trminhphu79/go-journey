package task

type TaskService interface {
	GetAllTasks() ([]Task, error)
	GetTaskById(id int) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(id int, task Task) (Task, error)
}
