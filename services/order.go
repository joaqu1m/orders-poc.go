package services

import (
	"orders-api/models"
	"orders-api/repositories"
)

func GetOrders() []models.Order {
	return repositories.GetOrders()
}
