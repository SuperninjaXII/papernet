// /api/index.go

package handler

import (
	"net/http"
	"search/config"
	"search/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Handler function expected by Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Initialize the Fiber engine
	engine := html.New("./../views", ".html")

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		Views: engine, // Set the template engine
	})

	// Serve static files
	app.Static("/", "./../public")

	// Define routes
	routes.Routes(app)

	// Handle HTTP requests with Fiber
	app.Listener = http.NewServeMux()
	app.Listener.ServeHTTP(w, r)
}

func init() {
	// Initialize the database before running the app
	config.Database()
}
