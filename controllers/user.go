package controllers

import (
	"net/http"
	"orders-api/dtos"
	"orders-api/services"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := services.GetUsers()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {

	var req dtos.UserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := services.CreateUser(
		dtos.UserRequest{
			Login:    req.Login,
			Password: req.Password,
		},
	)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

func AuthUser(c echo.Context) error {

	var req dtos.UserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	token, err := services.AuthUser(
		dtos.UserRequest{
			Login:    req.Login,
			Password: req.Password,
		},
	)
	if err != nil {
		print(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
