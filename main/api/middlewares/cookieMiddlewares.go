package middlewares

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func SetCookieMiddlewares(g *echo.Group) {
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.String(http.StatusUnauthorized, "You don't have any cookie")
			}
			log.Println(err)
			return err
		}
		if cookie.Value == "some_string" {
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "You don't have the right cookie")
	}
}
