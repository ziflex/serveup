package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/ziflex/serveup/internal/manifest"
)

type Endpoint struct {
	path    string
	methods []Method
}

func NewEndpoint(config manifest.Endpoint) (Endpoint, error) {
	methods := make([]Method, len(config.Methods))

	for i, m := range config.Methods {
		method, err := NewMethod(m)

		if err != nil {
			return Endpoint{}, errors.Wrapf(err, "create endpoint method: [%s] %s", m.Name, config.Path)
		}

		methods[i] = method
	}

	return Endpoint{
		path:    config.Path,
		methods: methods,
	}, nil
}

func (e *Endpoint) Path() string {
	return e.path
}

func (e *Endpoint) Use(router *echo.Echo) error {
	for _, method := range e.methods {
		switch method.Name() {
		case http.MethodGet:
			router.GET(e.path, method.Handle)
		case http.MethodPost:
			router.POST(e.path, method.Handle)
		case http.MethodPut:
			router.PUT(e.path, method.Handle)
		case http.MethodDelete:
			router.DELETE(e.path, method.Handle)
		case http.MethodPatch:
			router.PATCH(e.path, method.Handle)
		case http.MethodOptions:
			router.OPTIONS(e.path, method.Handle)
		case http.MethodHead:
			router.HEAD(e.path, method.Handle)
		case http.MethodConnect:
			router.CONNECT(e.path, method.Handle)
		case http.MethodTrace:
			router.TRACE(e.path, method.Handle)
		default:
			return errors.Errorf("invalid endpoint method name: [%s] %s", method.Name(), e.path)
		}
	}

	return nil
}
