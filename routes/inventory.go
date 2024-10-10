package routes

import (
	"github.com/Nivas-Mekala/IKEA_Inventory_Application/database"
	"github.com/Nivas-Mekala/IKEA_Inventory_Application/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func CreateInventory(c *fiber.Ctx) error {

	var inputRequest models.InventoryRequest
	var data models.Inventory

	if err := c.BodyParser(&inputRequest); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	for _, inv := range inputRequest.Inventory {
		data.Article_id = inv.Article_id
		data.Name = inv.Name
		data.Stock = inv.Stock
		database.Database.Db.Create(&data)
	}

	return c.Status(200).SendString("Inventory Data Saved Successfully!!")
}

func UpdateInventory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	response := models.Inventory{}

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findInventory(id, &response); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updateInv := models.Inventory{}
	if err := c.BodyParser(&updateInv); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.Db.Save(&updateInv)

	return c.Status(200).JSON(updateInv)

}

func findInventory(id int, inventory *models.Inventory) error {
	database.Database.Db.Find(&inventory, "article_id=?", id)
	return nil
}

func GetInventory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	response := models.Inventory{}

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findInventory(id, &response); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}

func GetAllInventory(c *fiber.Ctx) error {
	inventory := []models.Inventory{}
	database.Database.Db.Find(&inventory)
	return c.Status(200).JSON(inventory)
}

func DeleteInventory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var inventory models.Inventory

	if err := findInventory(id, &inventory); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Delete(&inventory)

	return c.Status(200).SendString("Inventory has been deleted")

}

func CreateInventoryResponse(inventory models.Inventory) models.Inventory {

	return models.Inventory{
		Article_id: inventory.Article_id,
		Name:       inventory.Name,
		Stock:      inventory.Stock,
	}
}
