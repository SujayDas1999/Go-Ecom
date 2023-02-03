package routes

import (
	"ecom/database"
	"ecom/helper"
	"ecom/models"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)



func CreateProduct(c *fiber.Ctx) error {
	
	var product models.Product

	log.Println(product)
	log.Println(c.BodyParser(product))

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseProduct := helper.CreateResponseProduct(&product)

	return c.Status(200).JSON(responseProduct)
}


func GetAllProducts(c *fiber.Ctx) error {
	
	products := []models.Product{}

	database.Database.Db.Find(&products)

	responseProducts := []helper.Product{}

	for _,product := range products {
		responseProduct := helper.CreateResponseProduct(&product)
		responseProducts = append(responseProducts, responseProduct) 
	}

	return c.Status(200).JSON(responseProducts)
}

func SearchBySerialNumber(c *fiber.Ctx) error {
	
	pSlno,err := c.ParamsInt("slno")

	if err != nil {
		return c.Status(400).JSON("Invalid slno")
	}

	var product models.Product

	database.Database.Db.Where("serial_number=?",pSlno).Find(&product)

	if product.ID == 0 {
		return c.Status(404).JSON("Product not found")
	}
 
	responseProduct := helper.CreateResponseProduct(&product)

	return c.Status(200).JSON(responseProduct)
}

func GetProductById(c *fiber.Ctx) error {
	
	prodId,err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid ID")
	}

	var product models.Product

	database.Database.Db.Where("id=?",prodId).Find(&product)

	if product.ID == 0 {
		return c.Status(404).JSON("Product not found")
	}

	responseProduct := helper.CreateResponseProduct(&product)

	return c.Status(200).JSON(responseProduct)

}

func UpdateProduct(c *fiber.Ctx) error {
	prodId,err := c.ParamsInt("id")
	
	if err != nil {
		return c.Status(400).JSON("Invalid ID")
	}

	var product models.Product

	database.Database.Db.Where("id=?",prodId).Find(&product)

	type UpdatedProduct struct {
		Name string `json:"name"`
		SerialNumber string `json:"slno"`
	}


	var updateProduct UpdatedProduct

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(400).JSON("Invalid Product Details")
	}

	var product2 models.Product

	database.Database.Db.Where("serial_number=?",updateProduct.SerialNumber).Find(&product2)

	if product2.ID == 0 {
		
		product.Name = updateProduct.Name
		product.SerialNumber = updateProduct.SerialNumber
	
	} else {
		return c.Status(400).JSON("SerialNumberExists")
	}

	database.Database.Db.Save(&product)
	
	responseProduct := helper.CreateResponseProduct(&product)

	return c.Status(200).JSON(responseProduct)

}

func DeleteProduct(c *fiber.Ctx) error {
	
	prodId,err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Bad ID")
	}

	var product models.Product

	database.Database.Db.Where("id=?",prodId).Delete(&product)
	
	responseProduct := helper.CreateResponseProduct(&product)

	return c.Status(200).JSON(responseProduct)

}


func findProduct(id int, product *models.Product ) error {
	database.Database.Db.Find(&product,"id=?",id)

	if product.ID == 0 {
		return errors.New("Product not found")
	}else {
		return nil
	}
	 
}