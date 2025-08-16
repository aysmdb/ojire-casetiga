// Package router contain all routes in app
package router

import (
	"github.com/gofiber/fiber/v2"
)

func ViewRoutes(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home/index", fiber.Map{}, "layouts/main")
	})
}
