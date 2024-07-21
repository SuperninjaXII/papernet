package routes

import (
	"papernet/controllers"

	"github.com/gofiber/fiber/v3"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.Home)
}
