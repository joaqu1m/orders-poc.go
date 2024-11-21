package routes

import (
	"github.com/labstack/echo/v4"
	"orders-api/controllers"
)

func OrderRoutes(e *echo.Echo) {
	e.GET("/orders", controllers.GetOrders)
}
