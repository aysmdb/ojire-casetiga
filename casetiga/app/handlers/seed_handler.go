// Package handlers for handling api request
package handlers

import (
	"github.com/aysmdb/ojire-casetiga/app/models"
	"github.com/gofiber/fiber/v2"
)

func SeedDB(c *fiber.Ctx) error {
	err := models.SeedProduct()
	if err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Failed to seed product",
		})
	}

	err = models.Seeduser()
	if err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Failed to seed user",
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Seeding DB success.",
	})
}
