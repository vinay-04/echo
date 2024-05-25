package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Yello(c echo.Context) error {
	return c.String(http.StatusOK, "Helo from server")
}

func Getval(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	profile := c.Param("profile")
	if profile == "admin" {
		return c.String(http.StatusOK, "Name: "+name+" Age: "+age+" Profile: "+profile)
	} else if profile == "user" {
		return c.String(http.StatusOK, "Name: "+name+" Age: "+age+"Profile: "+profile)
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Profile"})
}
