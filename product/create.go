package product

import (
	"github.com/labstack/echo"
	"net/http"
)

func List(c echo.Context) error {
	return c.String(http.StatusOK, "Listing products... :-)")
}

func Show(c echo.Context) error {
	out_str := "I should show a product with id: " + c.Param("id")
	// We can do Query strings too
	out_str += "\nSpecial querystring param: " + c.QueryParam("special")
	return c.String(http.StatusOK, out_str)
}

func Create(c echo.Context) error {
	// I should create a product with given params
	out_map := map[string]string{
		"isbn": c.FormValue("isbn"),
		"title": c.FormValue("title"),
	}
	return c.JSON(http.StatusOK, out_map)
}

// Middleware
func Trail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("A product create request was sent")
		return next(c)
	}
}
