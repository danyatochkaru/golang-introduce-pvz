package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pvz/models"
)

func GetAllStatuses(c echo.Context) error {
	var statuses []models.Status

	db.Find(&statuses)

	return c.JSON(http.StatusOK, statuses)
}

func GetStatusById(c echo.Context) error {
	var status models.Status

	db.Take(&status, c.Param("id"))

	if status.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, status)
}

func CreateStatus(c echo.Context) error {
	var statusValue models.StatusValue
	statusValue = models.StatusValue(c.FormValue("fullName"))

	order := models.Status{
		Value: statusValue,
	}

	db.Create(&order)

	return c.String(http.StatusCreated, "Order created")
}

func UpdateStatus(c echo.Context) error {
	var order models.Status

	db.Take(&order, c.Param("id"))

	if order.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	order.Value = models.StatusValue(c.FormValue("value"))

	db.Save(&order)

	return c.JSON(http.StatusOK, order)
}

func DeleteStatus(c echo.Context) error {
	var status models.Status

	db.Take(&status, c.Param("id"))

	if status.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	db.Delete(&status)

	return c.NoContent(http.StatusOK)
}
