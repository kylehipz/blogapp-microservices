package api

import (
	"github.com/labstack/echo/v4"
)

type EchoAPIRoute struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}
