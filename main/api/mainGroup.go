package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.Yello)
	e.GET("/login", handlers.Login)
	e.GET("/values/:profile", handlers.Getval)
	e.POST("/adduser", handlers.Adduser)
}
