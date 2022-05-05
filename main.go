package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	router "fazTudo-API/routes"
)

func main() {
	var e *echo.Echo = echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	println(e)

	router.LoadAllRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}
