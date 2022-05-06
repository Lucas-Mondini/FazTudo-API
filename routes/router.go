package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoadAllRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello World")
	})
	clientRoute(e)
}
