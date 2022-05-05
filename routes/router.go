package router

import "github.com/labstack/echo/v4"

func LoadAllRoutes(e *echo.Echo) {
	clientRoute(e)
}
