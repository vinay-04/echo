package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "ja ck" && password == "1234" {
		cookie := &http.Cookie{} // cookie := new(http.cookie)
		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)
		token, err := createToken()
		if err != nil {
			log.Println("Error creating token", err)
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}
		JWTCookie := &http.Cookie{} // cookie := new(http.cookie)
		JWTCookie.Name = "JWTCookie"
		JWTCookie.Value = token
		JWTCookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(JWTCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "Your username or password is incorrect")

}

func createToken() (string, error) {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id:        "admin_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err

	}
	return token, nil
}
