package router

import (
	"api"
	"api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	adminGroup := e.Group("/admin")

	cookieGroup := e.Group("/cookie")

	jwtGroup := e.Group("/jwt")

	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	// main routes
	api.MainGroup(e)
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JwtGroup(jwtGroup)

	return e
}
