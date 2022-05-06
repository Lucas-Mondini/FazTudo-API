package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"fazTudo-API/routes"
)

func main() {
	var e *echo.Echo = echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.Use(middleware.CORS())

	routes.LoadAllRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":9191"))
}
