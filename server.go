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

	e.PUT("/orders", handlers.UpdateOrder)

	e.DELETE("/orders", handlers.DeleteOrder)

	e.Logger.Fatal(e.Start(":1323"))
}
