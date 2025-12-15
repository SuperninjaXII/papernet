package main

import (
	"log"
	"papernet/config"
	"papernet/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.Database()

}

func main() {
	app := fiber.New(
		fiber.Config{
			BodyLimit: 524288000,
		},
	)
	app.Static("/", "./public")
	routes.Routes(app)
	routes.Admin(app)

	log.Fatal(app.Listen(":3000"))
}
