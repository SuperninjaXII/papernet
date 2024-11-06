package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Handler is the entry point for Vercel to call
func Handler(w http.ResponseWriter, r *http.Request) {
	// Create a new Fiber app
	app := fiber.New()

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber on Vercel!")
	})

	// Fiber app's handler that implements http.Handler
	handler := app
	// Use the Fiber app's handler (which is an HTTP handler)
	handler.ServeHTTP(w, r)
}
