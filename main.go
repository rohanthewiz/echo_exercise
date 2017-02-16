package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", rootHandler)
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, trackUsers)	

	e.Logger.Fatal(e.Start(":1323"))
}

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func trackUsers(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("A request is being sent to /users")
		return next(c)
	}
}
