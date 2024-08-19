package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pvz/models"
	"pvz/utils"
)

func GetAllProducts(c echo.Context) error {
	var products []models.Product

	db.Find(&products)

	return c.JSON(http.StatusOK, products)
}

func GetProductById(c echo.Context) error {
	var product models.Product

	db.Take(&product, c.Param("id"))

	if product.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	product := models.Product{
		Title: c.FormValue("title"),
		Price: utils.StringToUint(c.FormValue("price")),
	}

	db.Create(&product)

	return c.String(http.StatusCreated, "Product created")
}

func UpdateProduct(c echo.Context) error {
	var product models.Product

	db.Take(&product, c.Param("id"))

	if product.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	product.Title = c.FormValue("title")
	product.Price = utils.StringToUint(c.FormValue("price"))

	db.Save(&product)

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	var product models.Product

	db.Take(&product, c.Param("id"))

	if product.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	db.Delete(&product)

	return c.NoContent(http.StatusOK)
}
