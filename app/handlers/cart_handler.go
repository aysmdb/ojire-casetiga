package handlers

import (
	"github.com/aysmdb/ojire-casetiga/app/models"
	"github.com/gofiber/fiber/v2"
)

func AddToCartHandler(c *fiber.Ctx) error {
	r := new(models.CartRequest)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	cart := models.Cart{
		UserID:    1,
		ProductID: r.ProductID,
		Quantity:  r.Quantity,
	}

	err := models.AddToCart(cart)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add to cart: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product added to cart successfully",
	})
}

func GetCartByUserIDHandler(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	cartItems, err := models.GetCartByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve cart: " + err.Error(),
		})
	}

	return c.JSON(cartItems)
}
