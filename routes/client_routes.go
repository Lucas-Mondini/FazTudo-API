package routes

import (
	"fazTudo-API/view"

	"github.com/labstack/echo/v4"
)

func clientRoute(e *echo.Echo) {
	println(e)
	e.GET("/clients/:id", view.GetCustomer)
	e.GET("/clients", view.GetAllCustomers)
	e.POST("/clients", view.CreateCustomer)
	e.PUT("/clients", view.UpdateCustomer)
	e.DELETE("/clients", view.DeleteCustomer)
}
