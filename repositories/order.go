package repositories

import (
	"orders-api/database"
	"orders-api/models"
)

func GetOrders() []models.Order {

	rows, err := database.MySQLQuery("SELECT id, customer_name, customer_email, status, source FROM `order`")
	if err != nil {
		panic(err.Error())
	}

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(
			&order.ID,
			&order.CustomerName,
			&order.CustomerEmail,
			&order.Status,
			&order.Source,
		); err != nil {
			panic(err.Error())
		}
		orders = append(orders, order)
	}

	return orders
}
