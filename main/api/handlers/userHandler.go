package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Profile string `json:"profile"`
}

func Adduser(c echo.Context) error {
	user := User{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&user)

	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("User: %#v", user)
	return c.String(http.StatusOK, "We got your user!!!")
}
