package helper

import (
	"ecom/models"
	"time"
)

type Order struct {
	ID      uint    `json:"id"`
	Product Product `json:"product"`
	User    User    `json:"user"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order *models.Order, user User, product Product) Order {
	return Order {
		ID: order.ID,
		User: user,
		Product: product,
		CreatedAt: order.CreatedAt,
	}
}
