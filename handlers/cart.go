package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pvz/models"
	"pvz/utils"
)

func GetAllCarts(c echo.Context) error {
	var carts []models.Cart

	db.Find(&carts)

	return c.JSON(http.StatusOK, carts)
}

func GetCartById(c echo.Context) error {
	var cart models.Cart

	db.Preload("Products").Take(&cart, c.Param("id"))

	if cart.ID == 0 {
		return c.String(http.StatusNotFound, "Cart not found")
	}

	return c.JSON(http.StatusOK, cart)
}

func CreateCart(c echo.Context) error {
	cart := models.Cart{}

	db.Create(&cart)

	return c.JSON(http.StatusCreated, cart)
}

func DeleteCart(c echo.Context) error {
	var cart models.Cart

	db.Take(&cart, c.Param("id"))

	if cart.ID == 0 {
		return c.String(http.StatusNotFound, "Cart not found")
	}

	db.Delete(&cart)

	return c.NoContent(http.StatusOK)
}

func GetProductInCart(c echo.Context) error {
	var cart models.Cart

	db.Preload("Products").Find(&cart)

	return c.JSON(http.StatusOK, &cart.Products)
}

func AddProductToCart(c echo.Context) error {
	var cartProducts models.CartProducts

	db.Preload("Product").Preload("Order").Where(&models.CartProducts{
		CartID:    utils.StringToUint(c.Param("cart_id")),
		ProductID: utils.StringToUint(c.Param("product_id")),
	}).Take(&cartProducts)

	if cartProducts.ID != 0 {
		return c.String(http.StatusConflict, "Product already in cart")
	}

	cartProducts.CartID = utils.StringToUint(c.Param("cart_id"))
	cartProducts.ProductID = utils.StringToUint(c.Param("product_id"))
	db.Create(&cartProducts)

	var cart models.Cart

	db.Preload("Product").Preload("Order").Find(&cart)

	return c.JSON(http.StatusCreated, &cart)
}

func RemoveProductFromCart(c echo.Context) error {
	var cartProducts models.CartProducts

	db.Preload("Product").Preload("Order").Where(&models.CartProducts{
		CartID:    utils.StringToUint(c.Param("cart_id")),
		ProductID: utils.StringToUint(c.Param("product_id")),
	}).Take(&cartProducts)

	if cartProducts.ID == 0 {
		return c.String(http.StatusConflict, "Product not in cart yet")
	}

	db.Delete(&cartProducts)

	var cart models.Cart

	db.Preload("Product").Preload("Order").Find(&cart)

	return c.JSON(http.StatusOK, &cart)
}
