package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

type Server struct {
	version    string
	appVersion string
	appName    string
	port       uint64
	endpoints  []Endpoint
}

func New(settings Settings) (*Server, error) {
	endpoints := make([]Endpoint, len(settings.Manifest.Endpoints))

	for i, e := range settings.Manifest.Endpoints {
		endpoint, err := NewEndpoint(e)

		if err != nil {
			return nil, errors.Wrapf(err, "create endpoint: %s", e.Path)
		}

		endpoints[i] = endpoint
	}

	return &Server{
		version:    settings.Version,
		appVersion: settings.Manifest.Version,
		appName:    settings.Manifest.Name,
		port:       settings.Port,
		endpoints:  endpoints,
	}, nil
}

func (s *Server) Run() error {
	router := echo.New()
	router.HideBanner = true

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	router.Use(middleware.BodyLimit("1M"))
	router.Use(middleware.RequestID())
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/_/version", s.showVersion)
	router.GET("/_/health", s.healthCheck)

	for _, endpoint := range s.endpoints {
		if err := endpoint.Use(router); err != nil {
			return errors.Wrapf(err, "start endpoint: %s", endpoint.Path())
		}
	}

	router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	return router.Start(fmt.Sprintf("0.0.0.0:%d", s.port))
}

func (s *Server) showVersion(ctx echo.Context) error {
	payload := map[string]string{
		"server": s.version,
	}

	payload[s.appName] = s.appVersion

	return ctx.JSON(http.StatusOK, payload)
}

func (s *Server) healthCheck(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
