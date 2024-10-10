package main

import (
	"log"

	"github.com/Nivas-Mekala/IKEA_Inventory_Application/database"
	"github.com/Nivas-Mekala/IKEA_Inventory_Application/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectToDatabase()
	app := fiber.New()
	setUpRoutes(app)

	log.Fatal(app.Listen("127.0.0.1:8080"))
}

func setUpRoutes(app *fiber.App) {
	app.Get("/ikea", welcome)

	// Inventory Routes
	app.Post("/ikea/inventory/saveinventory", routes.CreateInventory)
	app.Get("/ikea/inventory", routes.GetAllInventory)
	app.Get("/ikea/inventory/:id", routes.GetInventory)
	app.Put("/ikea/inventory/:id", routes.UpdateInventory)
	app.Delete("/ikea/inventory/:id", routes.DeleteInventory)

	// Product Routes
	app.Post("/ikea/products/saveproducts", routes.CreateProducts)
	app.Get("/ikea/products", routes.GetAllProducts)
}

func welcome(c *fiber.Ctx) error {
	return c.Status(200).SendString("Welcome to IKEA - To Create better everyday life!!")
}
