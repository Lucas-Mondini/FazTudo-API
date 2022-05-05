package router

import (
	clientView "fazTudo-API/view/Client"

	"github.com/labstack/echo/v4"
)

func clientRoute(e *echo.Echo) {
	println(e)
	e.GET("Clients/:id", clientView.GetClient)
}
