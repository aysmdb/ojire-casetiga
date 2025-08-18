package handlers

import (
	"github.com/aysmdb/ojire-casetiga/app/models"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	take:= c.QueryInt("take", 9)
	search := c.Query("search", "")

	products := models.GetProductList(page, take, search)

	return c.JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}
	
	product := models.GetProductByID(id)

	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return c.JSON(product)
}
