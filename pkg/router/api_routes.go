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
}
