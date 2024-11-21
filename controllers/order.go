package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"orders-api/services"
)

func GetOrders(c echo.Context) error {
	orders := services.GetOrders()

	return c.JSON(http.StatusOK, orders)
}
