package routes

import (
	"papernet/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Post("/result", controllers.SearchResult)

	app.Get("/books", controllers.GetAllBooks)
	app.Get("/book/:id", controllers.GetBookByID)
	app.Post("/search", controllers.SearchBooks)

}
