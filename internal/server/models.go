package server

import "github.com/jerrdasur/someTodo/internal/todo"

type ErrorResponse struct {
	Error string `json:"error"`
}

type TodoListResponse struct {
	Data []Todo `json:"data"`
}

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func NewTodoListResponse(l []todo.Task) TodoListResponse {
	data := make([]Todo, len(l))
	for i := range l {
		data[i] = NewTodo(l[i])
	}

	return TodoListResponse{
		Data: data,
	}
}

func NewTodo(t todo.Task) Todo {
	return Todo{
		ID:    t.ID,
		Title: t.Title,
	}
}

type CreateTodoRequest struct {
	Title string `json:"title"`
}
