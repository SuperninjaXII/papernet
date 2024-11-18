package routes

import (
	"papernet/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	//html routes
	app.Get("/", controllers.Home)
	app.Get("/admin", controllers.Admin)
	app.Get("/search", controllers.SearchPage)
	app.Post("/SearchResults", controllers.BooksPage)
	app.Post("/result", controllers.SearchResult)
	app.Get("/download/:id", controllers.DownloadPage)
	app.Get("/related/:tag", controllers.GetBookByCartegoryBooks)
	//json
	app.Get("/books", controllers.GetAllBooks)
	app.Get("/book/:id", controllers.GetBookByID)
	app.Post("/search", controllers.SearchBooks)
	app.Get("/category/:tag", controllers.GetBookByCartegory)
}
