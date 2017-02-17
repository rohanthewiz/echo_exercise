package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/rohanthewiz/echo_exercise/product"
)

func main() {
	e := echo.New()

	e.GET("/", root)

	e.GET("/ping", func(c echo.Context) error {  // we can use an anonymous function
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/products", product.List)
	e.GET("/product/:id", product.Show)
	e.POST("/products", product.Create, product.Trail)

	// Start listening
	e.Logger.Fatal(e.Start(":1323"))
}

func root(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
