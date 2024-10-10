package routes

import (
	"github.com/Nivas-Mekala/IKEA_Inventory_Application/database"
	"github.com/Nivas-Mekala/IKEA_Inventory_Application/internal/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProducts(c *fiber.Ctx) error {
	var inputRequest models.ProductsRequest
	var data models.Product

	if err := c.BodyParser(&inputRequest); err != nil {
		c.Status(400).JSON(err.Error())
	}

	for _, inv := range inputRequest.Products {
		data.Product_name = inv.Name
		for _, value := range inv.ContainArticles {
			data.Article_id = value.ArtID
			data.Amount_of = value.AmountOf
			database.Database.Db.Create(&data)
		}
	}

	return c.Status(200).SendString("Products Data Saved Successfully!!")
}

func GetAllProducts(c *fiber.Ctx) error {

	products := []models.Product{}
	database.Database.Db.Find(&products)

	for _, product := range products {
		product := CreateResponseProducts(product)
		products = append(products, product)
	}

	return c.Status(200).JSON(products)
}

func CreateResponseProducts(userModel models.Product) models.Product {
	return models.Product{
		Product_name: userModel.Product_name,
		Article_id:   userModel.Article_id,
		Amount_of:    userModel.Amount_of,
	}
}
