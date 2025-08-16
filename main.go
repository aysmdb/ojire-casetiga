package main

import (
	"github.com/aysmdb/ojire-casetiga/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.ViewRoutes(app)

	app.Listen(":3200")
}
