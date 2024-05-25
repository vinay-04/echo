package middlewares

import (
	"github.com/labstack/echo"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(serverHeader)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "CustomVR/1.0")
		return next(c)
	}
}
