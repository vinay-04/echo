package api

import (
	"api/handlers"

	"github.com/labstack/echo/v4"
)

func AdminGroup(g *echo.Group) {
	g.Group("/admin", handlers.MainAdmin)
}
