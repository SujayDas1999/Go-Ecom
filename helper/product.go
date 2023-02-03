package helper

import "ecom/models"

type Product struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	SerialNumber string `json:"slno"`
}


func CreateResponseProduct(product *models.Product) Product {
	return Product {
		ID: product.ID,
		Name: product.Name,
		SerialNumber: product.SerialNumber,
	}
}