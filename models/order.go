package models

type Order struct {
	ID            int64  `json:"id"`
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
	Status        string `json:"status"`
	Source        string `json:"source"`
}
