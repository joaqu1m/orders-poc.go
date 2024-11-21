package main

import (
	"log"
	"orders-api/config"
	"orders-api/routes"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	e := echo.New()
	e.Use(middleware.Recover())

	e.Use(echojwt.WithConfig(config.GetJWTConfig()))

	routes.OrderRoutes(e)
	routes.UserRoutes(e)

	log.Printf("Server running on port %s", os.Getenv("SERVER_PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
