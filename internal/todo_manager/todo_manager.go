package todo_manager

import (
	"fmt"

	"github.com/sharmajidotdev/go_todo/internal/todo"
)

type TodoManager struct {
	Todos map[int]*todo.Todo
}

func New() *TodoManager {
	return &TodoManager{
		Todos: make(map[int]*todo.Todo),
	}
}

func (manager *TodoManager) AddTodo(id int, description string) {
	todo := todo.New(id, description)
	manager.Todos[id] = todo
}

func (manager *TodoManager) DeleteTodo(id int) {
	if _, exists := manager.Todos[id]; exists {
		delete(manager.Todos, id)
		fmt.Printf("Todo ID %d deleted.\n", id)
	} else {
		fmt.Printf("Todo with ID %d not found.\n", id)
	}
}

func (manager *TodoManager) DeleteAllTodo() {
	if len(manager.Todos) != 0 {
		for id := range manager.Todos {
			delete(manager.Todos, id)
			fmt.Printf("Todo ID %d deleted.\n", id)
		}
		fmt.Printf("All deleted.\n")
	} else {
		fmt.Printf("Already empty.\n")
	}
}

func (manager *TodoManager) UpdateTodo(id int, description string) {
	if todo, exists := manager.Todos[id]; exists {
		todo.SetDescription(description)
	} else {
		fmt.Printf("Todo with ID %d not found.\n", id)
	}
}

func (manager *TodoManager) DisplayTodos() {
	for _, todo := range manager.Todos {
		todo.DisplayTodo()
	}
}
