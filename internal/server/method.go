package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo"

	"github.com/ziflex/serveup/internal/manifest"
	"github.com/ziflex/serveup/internal/runtime"
)

type Method struct {
	name        string
	description string
	program     runtime.Program
}

func NewMethod(config manifest.Method) (Method, error) {
	program, err := runtime.NewProgram(config.Program)

	if err != nil {
		return Method{}, err
	}

	return Method{
		name:        config.Name,
		description: config.Description,
		program:     program,
	}, nil
}

func (m *Method) Name() string {
	return m.name
}

func (m *Method) Description() string {
	return m.description
}

func (m *Method) Handle(ctx echo.Context) error {
	res, err := m.program.Exec(
		context.WithValue(ctx.Request().Context(), "req", NewRequestContext(ctx)),
	)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, Failure{err})
	}

	return ctx.JSON(http.StatusOK, Success{string(res)})
}
