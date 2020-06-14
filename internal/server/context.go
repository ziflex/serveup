package server

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
)

type RequestContext struct {
	echo echo.Context

	URL *url.URL
}

func NewRequestContext(echo echo.Context) RequestContext {
	return RequestContext{
		echo: echo,
		URL:  echo.Request().URL,
	}
}

func (rc *RequestContext) Param(name string) string {
	return rc.echo.Param(name)
}

func (rc *RequestContext) Query(name string) string {
	return rc.echo.QueryParam(name)
}

func (rc *RequestContext) Cookie(name string) (*http.Cookie, error) {
	return rc.echo.Cookie(name)
}

func (rc *RequestContext) Header(name string) string {
	return rc.echo.Request().Header.Get(name)
}

func (rc *RequestContext) Body() (string, error) {
	out, err := ioutil.ReadAll(rc.echo.Request().Body)

	if err != nil {
		return "", err
	}

	return string(out), nil
}
