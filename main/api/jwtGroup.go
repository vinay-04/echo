package api

import (
	"api/handlers"

	"github.com/labstack/echo/v4"
)

func JwtGroup(g *echo.Group) {
	g.Group("/jwt", handlers.MainJwt)
}
