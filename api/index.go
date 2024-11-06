package handler

import (
	"log"
	"net/http"
	"search/config"
	"search/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Log to stdout (will be visible in Vercel logs)
	log.Println("Starting Fiber app...")

	// Initialize the Fiber engine
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

	// Handle HTTP requests with Fiber
	app.Listener = http.NewServeMux()
	app.Listener.ServeHTTP(w, r)
}

func init() {
	// Initialize the database before running the app
	log.Println("Initializing database...")
	config.Database()
}
