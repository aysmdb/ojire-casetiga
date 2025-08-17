package router

import (
	"github.com/aysmdb/ojire-casetiga/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func APIRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/seed", handlers.SeedDB)

	route.Get("/product/list", handlers.GetProducts)
	route.Get("/product/:id", handlers.GetProductByID)

	route.Post("/user/login", handlers.LoginHandler)
	route.Get("/user/:id", handlers.GetUserByIDHandler)

	route.Post("/cart/add", handlers.AddToCartHandler)
	route.Get("/cart/:user_id", handlers.GetCartByUserIDHandler)
}
