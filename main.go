package main

import (
	"ecom/database"
	"ecom/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New();
	
	setupRoutes(app)

	app.Listen(":3000")
}


func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)
	
	//USER ROUTES RELATED
	app.Post("/api/users/create",routes.CreateUser)
	app.Get("/api/users/all",routes.GetUsers)
	app.Get("/api/users/:id",routes.GetUserById)	
	app.Put("/api/users/:id",routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	//PRODUCT ROUTES RELATED
	app.Post("/api/products/create",routes.CreateProduct)
	app.Get("/api/products/all",routes.GetAllProducts)
	app.Get("/api/products/:slno",routes.SearchBySerialNumber)
	app.Get("/api/products/get/:id",routes.GetProductById)
	app.Put("/api/products/update/:id",routes.UpdateProduct)
	app.Delete("/api/products/delete/:id",routes.DeleteProduct)

	//ORDER ROUTES RELATED
	app.Post("/api/orders/create",routes.CreateOrder)
	app.Get("/api/orders/all",routes.GetAllOrders)


}


func welcome(c *fiber.Ctx) error {
	return c.JSON("mike")
}