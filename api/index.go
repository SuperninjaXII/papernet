package main

import (
	"log"
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

	// Convert the Fiber app into an HTTP handler
	http.Handle("/", app)

	// Call the app handler
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
