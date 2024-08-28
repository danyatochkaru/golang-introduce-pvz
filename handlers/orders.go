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

	db.Preload("Status").Preload("Products").Take(&order, c.Param("id"))

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

func SetOrderStatus(c echo.Context) error {
	var (
		order  models.Order
		status models.Status
	)

	db.Take(&order, c.Param("id"))
	db.Take(&status, c.FormValue("statusId"))

	if order.ID == 0 {
		return c.String(http.StatusNotFound, "Order not found")
	}
	if status.ID == 0 {
		return c.String(http.StatusBadRequest, "Status not found")
	}

	order.Status = status

	db.Save(&order)

	return c.JSON(http.StatusOK, order)
}

func AddProductToOrder(c echo.Context) error {
	var (
		order   models.Order
		product models.Product
	)

	db.Take(&order, c.Param("id"))
	db.Take(&product, c.FormValue("productId"))

	if order.ID == 0 {
		return c.String(http.StatusNotFound, "Order not found")
	}
	if product.ID == 0 {
		return c.String(http.StatusBadRequest, "Product not found")
	}

	order.Products = append(order.Products, &product)

	db.Save(&order)

	return c.JSON(http.StatusOK, order)
}
