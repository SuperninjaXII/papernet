package routes

import (
	"papernet/controllers"

	"github.com/gofiber/fiber/v2"
)

// this for routes hyou can use without the authorization header
func Routes(app *fiber.App) {
	app.Get("/books", controllers.GetAllBooks)
}
