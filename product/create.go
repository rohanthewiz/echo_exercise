package product

import (
	"github.com/labstack/echo"
	"net/http"
)

func Create(c echo.Context) error {
	return c.String(http.StatusOK, "Yeah I created that :-)")
}

func List(c echo.Context) error {
	return c.String(http.StatusOK, "Listing products... :-)")
}

// Middleware
func Trail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("A product create request was sent")
		return next(c)
	}
}
