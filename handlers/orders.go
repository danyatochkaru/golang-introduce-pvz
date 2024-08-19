package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pvz/models"
)

func GetAllOrders(c echo.Context) error {
	var orders []models.Order

	db.Find(&orders)

	return c.JSON(http.StatusOK, orders)
}

func GetOrderById(c echo.Context) error {
	var order models.Order

	db.Take(&order, c.Param("id"))

	if order.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, order)
}

func CreateOrder(c echo.Context) error {
	order := models.Order{
		FullName: c.FormValue("fullName"),
	}

	db.Create(&order)

	return c.String(http.StatusCreated, "Order created")
}

func UpdateOrder(c echo.Context) error {
	var order models.Order

	db.Take(&order, c.Param("id"))

	if order.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	order.FullName = c.FormValue("fullName")

	db.Save(&order)

	return c.JSON(http.StatusOK, order)
}

func DeleteOrder(c echo.Context) error {
	var order models.Order

	db.Take(&order, c.Param("id"))

	if order.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	db.Delete(&order)

	return c.NoContent(http.StatusOK)
}
