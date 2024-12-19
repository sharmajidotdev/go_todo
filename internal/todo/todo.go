package todo

import "fmt"

type Todo struct {
	TodoId          int
	TodoDescription string
}

func (t *Todo) DisplayTodo() {
	fmt.Printf("Todo ID: %d\nDescription: %s\n", t.TodoId, t.TodoDescription)
}

func (t *Todo) SetDescription(description string) {
	t.TodoDescription = description
}

func New(id int, description string) *Todo {
	if description == "" {
		description = "na"
	}
	return &Todo{
		TodoId:          id,
		TodoDescription: description,
	}
}
