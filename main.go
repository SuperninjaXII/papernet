package main

import (
	"papernet/config"
	"papernet/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
)

func init() {
	config.Database()
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	routes.Routes(app)
	app.Listen(":3000")
}
