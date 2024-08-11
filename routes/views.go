package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Views struct {
}

func (r *Views) Setup(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "")
	})

	e.GET("/admin", func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin", "")
	})

}
