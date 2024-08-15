package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pvz/database"
	"pvz/models"
	"pvz/utils"
)

var db = database.GetDBConnection()

func GetAllOrders(c echo.Context) error {
	var orders []models.Order

	db.Find(&orders)

	return c.JSON(http.StatusOK, orders)
}

func GetOrderById(c echo.Context) error {
	var order models.Order

	db.Last(&order, c.Param("id"))

	return c.JSON(http.StatusOK, order)
}

func CreateOrder(c echo.Context) error {
	fullName := c.FormValue("fullName")
	id := utils.StringToUint(c.FormValue("id"))

	order := models.Order{FullName: fullName, ID: id}

	db.Create(&order)

	return c.String(http.StatusCreated, "Order created")
}

func UpdateOrder(c echo.Context) error {
	var order models.Order

	db.First(&order, c.Param("id"))

	if order.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	order.FullName = c.FormValue("fullName")

	db.Save(&order)

	return c.JSON(http.StatusOK, order)
}

func DeleteOrder(c echo.Context) error {
	var order models.Order

	db.First(&order, c.Param("id"))

	if order.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	db.Delete(&order)

	return c.NoContent(http.StatusOK)
}
