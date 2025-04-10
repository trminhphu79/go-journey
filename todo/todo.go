package TodoModule

import "fmt"

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoService interface {
	GetAllTodos() ([]Todo, error)
	GetTodoById(id int) (Todo, error)
	CreateTodo(todo Todo) (Todo, error)
	UpdateTodo(id int, todo Todo) (Todo, error)
	DeleteTodo(id int) (string, error)
}

type TodoServiceImpl struct {
	Todos []Todo `json:"todos"`
}

func (t *TodoServiceImpl) GetAllTodos() ([]Todo, error) {
	return t.Todos, nil
}

func (t *TodoServiceImpl) GetTodoById(id int) (Todo, error) {
	for _, todo := range t.Todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return Todo{}, fmt.Errorf("todo not found")
}

func (t *TodoServiceImpl) CreateTodo(todo Todo) (Todo, error) {
	t.Todos = append(t.Todos, todo)
	return todo, nil
}

func (t *TodoServiceImpl) UpdateTodo(id int, todo Todo) (Todo, error) {
	for i, todoItem := range t.Todos {
		if todoItem.ID == id {
			t.Todos[i] = todo
			return todo, nil
		}
	}
	return Todo{}, fmt.Errorf("todo not found")
}

func (t *TodoServiceImpl) DeleteTodo(id int) (string, error) {
	for i, todoItem := range t.Todos {
		if todoItem.ID == id {
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
			return fmt.Sprintf("Todo deleted with id: %d", id), nil
		}
	}
	return "", fmt.Errorf("todo not found")
}
