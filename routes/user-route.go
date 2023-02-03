package routes

import (
	"ecom/database"
	"ecom/helper"
	"ecom/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)






func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := helper.CreateResponseUser(&user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)

	responseUsers := []helper.User{}
	for _, user := range users {
		responseUser := helper.CreateResponseUser(&user)
		responseUsers = append(responseUsers, responseUser)
	} 

	return c.Status(200).JSON(responseUsers)

}

func GetUserById(c *fiber.Ctx) error {
	

	userId,err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Invalid ID")
	}

	if err := findUser(userId, &user); err != nil {
		return c.Status(404).JSON("User not found")
	}	


	responseUser := helper.CreateResponseUser(&user)

	return c.Status(200).JSON(responseUser)
}


func UpdateUser(c *fiber.Ctx) error {
	userId,err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid ID")
	}

	var user models.User

	if err := findUser(userId,&user); err != nil {
		return c.Status(404).JSON("User does not exists")
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`	
		LastName string `json:"last_name"`
	}

	var userUpdate UpdateUser

	if err := c.BodyParser(&userUpdate); err != nil {
		c.Status(400).JSON("Something wrong with user detail")
	}

	user.FirstName = userUpdate.FirstName
	user.LastName = userUpdate.LastName

	database.Database.Db.Save(&user)

	responseUser := helper.CreateResponseUser(&user)
	
	return c.Status(200).JSON(responseUser)
}


func DeleteUser(c *fiber.Ctx) error {
	userId,err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Invalid ID")
	}

	var user models.User

	if err := findUser(userId,&user); err != nil {
		return c.Status(404).JSON("User not found")
	}

	database.Database.Db.Delete(&user)

	responseUser := helper.CreateResponseUser(&user) 

	return c.Status(200).JSON(responseUser)

}



func findUser(id int, user *models.User ) error {
	database.Database.Db.Find(&user,"id=?",id)
	
	if user.ID == 0 {
		return errors.New("User does not exists")
	} 

	return nil
}


