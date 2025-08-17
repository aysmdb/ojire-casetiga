package main

import (
	"github.com/aysmdb/ojire-casetiga/app/models"
	"github.com/aysmdb/ojire-casetiga/pkg/database"
	"github.com/aysmdb/ojire-casetiga/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectSqlite() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("mancing.db"))
	if err != nil {
		panic("failed to connect to the database")
	}

	database.DBConn.AutoMigrate(&models.Product{})
	database.DBConn.AutoMigrate(&models.User{})
	database.DBConn.AutoMigrate(&models.Cart{})
}

func main() {
	connectSqlite()

	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.ViewRoutes(app)
	router.APIRoutes(app)

	app.Listen(":3200")
}
