package routes

import (
	"ecom/database"
	"ecom/helper"
	"ecom/models"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	} 

	var user models.User

	if err := findUser(order.UserRef, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product 

	if err := findProduct(order.ProductRef, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := helper.CreateResponseUser(&user)

	responseProduct := helper.CreateResponseProduct(&product)

	responseOrder := helper.CreateResponseOrder(&order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}


func GetAllOrders(c *fiber.Ctx) error {
	
	orders := []models.Order{}

	database.Database.Db.Find(&orders)
	responseOrders := []helper.Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		database.Database.Db.Find(&user,"id=?",order.UserRef)
		database.Database.Db.Find(&product,"id=?",order.ProductRef)

		responseUser := helper.CreateResponseUser(&user)
		responseProduct := helper.CreateResponseProduct(&product)

		resposeOrder := helper.CreateResponseOrder(&order, responseUser, responseProduct)

		responseOrders = append(responseOrders, resposeOrder)

	}

	return c.Status(200).JSON(responseOrders)

}