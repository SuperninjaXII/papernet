package routes

import (
	"papernet/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Get("/search", controllers.SearchPage)
	app.Post("/SearchResults", controllers.BooksPage)
	app.Post("/result", controllers.SearchResult)
	app.Get("/download/:id",controllers.DownloadPage)

	app.Get("/books", controllers.GetAllBooks)
	app.Get("/book/:id", controllers.GetBookByID)
	app.Post("/search", controllers.SearchBooks)

}
