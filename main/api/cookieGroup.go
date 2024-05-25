package api

import (
	"api/handlers" // Replace FILEPATH with the actual file path of the handlers package

	"github.com/labstack/echo/v4"
)

func CookieGroup(g *echo.Group) {
	g.Group("/cookie", handlers.MainCookie)
}
