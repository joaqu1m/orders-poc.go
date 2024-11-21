package routes

import (
	"github.com/labstack/echo/v4"
	"orders-api/controllers"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/users", controllers.GetUsers)
	e.POST("/user", controllers.CreateUser)
	e.POST("/auth", controllers.AuthUser)
}
