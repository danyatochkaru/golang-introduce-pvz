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
	status := models.Status{
		Value: c.FormValue("statusValue"),
	}

	db.Create(&status)

	return c.String(http.StatusCreated, "Order created")
}

func UpdateStatus(c echo.Context) error {
	var status models.Status

	db.Take(&status, c.Param("id"))

	if status.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	status.Value = c.FormValue("value")

	db.Save(&status)

	return c.JSON(http.StatusOK, status)
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
