package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", rootHandler)
	// Just to show we can use an anonymous function for handling
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/users", usersRoot, trackUsers)	

	// Start listening
	e.Logger.Fatal(e.Start(":1323"))
}


func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func usersRoot(c echo.Context) error {
	return c.String(http.StatusOK, "We would return a listing of users here")
}

func trackUsers(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("A request is being sent to /users")
		return next(c)
	}
}
