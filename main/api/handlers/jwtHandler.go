package handlers

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func MainJwt(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])
	return c.String(http.StatusOK, "You are on jwt page")
}
