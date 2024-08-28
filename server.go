package main

import (
	"github.com/labstack/echo/v4"
	"pvz/handlers"
)

func main() {
	e := echo.New()

	e.GET("/orders", handlers.GetAllOrders)
	e.GET("/orders/:id", handlers.GetOrderById)
	e.POST("/orders", handlers.CreateOrder)
	e.PUT("/orders/:id", handlers.UpdateOrder)
	e.DELETE("/orders/:id", handlers.DeleteOrder)

	e.PUT("/orders/:id/status", handlers.SetOrderStatus)

	e.GET("/products", handlers.GetAllProducts)
	e.GET("/products/:id", handlers.GetProductById)
	e.POST("/products", handlers.CreateProduct)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.DELETE("/products/:id", handlers.DeleteProduct)

	e.GET("/statuses", handlers.GetAllStatuses)
	e.GET("/statuses/:id", handlers.GetStatusById)
	e.POST("/statuses", handlers.CreateStatus)
	e.PUT("/statuses/:id", handlers.UpdateStatus)
	e.DELETE("/statuses/:id", handlers.DeleteStatus)

	e.GET("/carts", handlers.GetAllCarts)
	e.GET("/carts/:id", handlers.GetCartById)
	e.POST("/carts", handlers.CreateCart)
	e.DELETE("/carts/:id", handlers.DeleteCart)

	e.GET("/carts/:cart_id/products", handlers.GetProductsInCart)
	e.DELETE("/carts/:cart_id/products", handlers.ClearProductFromCart)
	e.POST("/carts/:cart_id/products/:product_id", handlers.AddProductToCart)
	e.PUT("/carts/:cart_id/products/:product_id", handlers.ChangeProductAmount)
	e.DELETE("/carts/:cart_id/products/:product_id", handlers.RemoveProductFromCart)

	e.Logger.Fatal(e.Start(":1323"))
}
