package clientView

import (
	"net/http"

	clientController "fazTudo-API/controller/Client"

	"github.com/labstack/echo/v4"
)

func GetClient(c echo.Context) error {

	param := c.Param("id")
	var re map[string]string
	if param != "" {
		re = clientController.GetClient(param)
	} else {
		re = clientController.GetAllClients()
	}

	return c.JSON(http.StatusOK, re)
}
