package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fazTudo-API/controller"

	"github.com/labstack/echo/v4"
)

//DONE
func GetAllCustomers(c echo.Context) error {

	customer := new(controller.DefaultCustomerRepository)

	return c.JSON(http.StatusOK, customer.FindAll())
}

//DONE
func GetCustomer(c echo.Context) error {

	param, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"error": "something went wrong",
		})
	}

	customer := new(controller.DefaultCustomerRepository)

	return c.JSON(http.StatusOK, customer.Find(param))
}

//DONE
func CreateCustomer(c echo.Context) error {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "something went wrong",
		})
	}

	user_name := json_map["Name"]

	customer := new(controller.DefaultCustomerRepository)

	return c.JSON(http.StatusOK, customer.Create(fmt.Sprintf("%v", user_name)))
}

func UpdateCustomer(c echo.Context) error {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}

	user_name := json_map["name"]
	user_id := json_map["id"]

	customer := new(controller.DefaultCustomerRepository)
	customer.Update(user_id.(int), user_name.(string))

	return c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c echo.Context) error {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}
	user_id := json_map["id"]
	customer := new(controller.DefaultCustomerRepository)

	return c.JSON(http.StatusOK, customer.Delete(user_id.(int)))
}
