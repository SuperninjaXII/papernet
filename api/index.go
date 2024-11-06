package handler

import (
	"fmt"
	"log"
	"net/http"
	"search/config"
	"search/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Log the start of the application
	log.Println("Starting Fiber app...")

	// Initialize the Fiber engine with templates
	engine := html.New("./views", ".html")

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		Views: engine, // Set the template engine
	})

	// Serve static files
	app.Static("/", "./public")

	// Define routes
	routes.Routes(app)

	// Log the fact that we are listening on a port
	log.Println("App is listening on :3000")

	// Start the Fiber app (automatically listens and serves)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func init() {
	// Initialize the database before running the app
	log.Println("Initializing database...")
	config.Database()
}
