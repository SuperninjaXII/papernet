package main

import (
	"log"
	"search/config"
	"search/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	// Initialize the database connection
	err := config.Database()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
}

func FiberHandler(c *fiber.Ctx) error {
	// Initialize the view engine for HTML templates
	engine := html.New("./views", ".html")

	// Create a new Fiber app with the views engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files from the public directory
	app.Static("/", "./public")

	// Set up routes for your application
	routes.Routes(app)

	// Return the server response to Vercel (using the handler)
	return app.Listen(":3000")
}

// This function is required for Vercel serverless deployment
func Handler(c *fiber.Ctx) error {
	// Call the FiberHandler function to handle the request
	return FiberHandler(c)
}

