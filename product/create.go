package product

import (
	"github.com/labstack/echo"
	"net/http"
)

func List(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Listing products... :-)")
}

func Show(ctx echo.Context) error {
	out_str := "I should show a product with id: " + ctx.Param("id")
	// We can do Query strings too
	out_str += "\nSpecial querystring param: " + ctx.QueryParam("special")
	return ctx.String(http.StatusOK, out_str)
}

func Create(ctx echo.Context) error {
	out_str := "I should create a product with params\n"
	out_str += "ISBN: " + ctx.FormValue("isbn")
	out_str += ", Title: " + ctx.FormValue("title") + "\n"
	return ctx.String(http.StatusOK, out_str)
}

// Middleware
func Trail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		println("A product create request was sent")
		return next(ctx)
	}
}
