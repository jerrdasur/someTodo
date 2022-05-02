package main

import (
	"github.com/labstack/gommon/log"

	"github.com/jerrdasur/someTodo/internal/server"
	"github.com/jerrdasur/someTodo/internal/todo"
)

func main() {
	s := server.New(server.Opts{
		TodoHandler: server.NewTodoHandler(todo.NewList()),
		Addr:        "localhost:9090",
	})

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
