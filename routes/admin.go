package routes

import (
	"papernet/controllers"

	"github.com/gofiber/fiber/v2"
)

func Admin(app *fiber.App) {
	app.Get("admin/AllData", controllers.DisplayAllBooksJson)
	app.Post("/admin/create", controllers.Create)
}
