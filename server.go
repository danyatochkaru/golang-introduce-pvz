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

	e.Logger.Fatal(e.Start(":1323"))
}
