package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jerrdasur/someTodo/internal/todo"
)

const (
	idParamName = "id"
)

type TodoListStorage interface {
	List() []todo.Task
	Get(todoID string) *todo.Task
	Add(title string) todo.Task
	Remove(todoID string) error
	Update(t todo.Task) error
}

type TodoHandler struct {
	tls TodoListStorage
}

func NewTodoHandler(tls TodoListStorage) *TodoHandler {
	return &TodoHandler{tls: tls}
}

func (th *TodoHandler) InitRoutes(g *echo.Group) {
	g.GET("/", th.List)
	g.GET("/:id", th.Get)
	g.POST("/", th.Add)
	g.PUT("/:id", th.Update)
	g.DELETE("/:id", th.Delete)
}

func (th *TodoHandler) List(c echo.Context) error {
	l := th.tls.List()

	return c.JSON(http.StatusOK, NewTodoListResponse(l))
}

func (th *TodoHandler) Add(c echo.Context) error {
	req := new(CreateTodoRequest)

	if err := c.Bind(req); err != nil {
		return fmt.Errorf("Add: %w", err)
	}

	t := th.tls.Add(req.Title)

	return c.JSON(http.StatusCreated, NewTodo(t))
}

func (th *TodoHandler) Delete(c echo.Context) error {
	id := c.Param(idParamName)

	if err := th.tls.Remove(id); errors.Is(err, todo.ErrTodoNotFound) {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Error: fmt.Sprintf("todo '%s' not found", id),
		})
	} else if err != nil {
		return fmt.Errorf("Delete: %w", err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (th *TodoHandler) Update(c echo.Context) error {
	return nil
}

func (th *TodoHandler) Get(c echo.Context) error {
	return nil
}
