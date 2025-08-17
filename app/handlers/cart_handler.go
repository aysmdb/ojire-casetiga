package handlers

import (
	"github.com/aysmdb/ojire-casetiga/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddToCartHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(uint)

	r := new(models.CartRequest)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	cart := models.Cart{
		UserID:    userID,
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
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(uint)

	cartItems, err := models.GetCartByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve cart: " + err.Error(),
		})
	}

	return c.JSON(cartItems)
}

func CheckoutHandler(c *fiber.Ctx) error {
	var checkoutRequest models.CheckoutRequest
	if err := c.BodyParser(&checkoutRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := models.UpdateProductQuantity(checkoutRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Checkout failed: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Checkout successful",
	})
}
