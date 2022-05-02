package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	todoHandler *TodoHandler
	addr        string
}

type Opts struct {
	TodoHandler *TodoHandler
	Addr        string
}

func New(opts Opts) *Server {
	return &Server{
		todoHandler: opts.TodoHandler,
		addr:        opts.Addr,
	}
}

func (s *Server) Run() error {
	e := echo.New()
	apiGroup := e.Group("/api")

	{
		v1 := apiGroup.Group("/v1")
		s.todoHandler.InitRoutes(v1.Group("/todos"))
	}

	return e.Start(s.addr)
}
