package routes

import (
    "github.com/gofiber/fiber/v2"
    "search/controllers"
    "search/config"
    "search/models"
)

func Routes(app *fiber.App) {
    app.Get("/books", controllers.GetAllBooks)
    app.Get("/books/:id", controllers.GetBookByID)
    app.Post("/upload", controllers.UploadHandler)
    app.Delete("/delete/:id", controllers.DeleteMsgHandler)

    app.Static("/", "./public")
    app.Get("/", func(c *fiber.Ctx) error {
        // Render index template
        return c.Render("Layout", fiber.Map{
            "Title": "PAPERNET",
        })
    })
    app.Get("/addBook", func(c *fiber.Ctx) error {
        // Render index template
        return c.Render("Create", fiber.Map{
            "user": "jose",
        })
    })
    app.Get("/download/:id", func(c *fiber.Ctx) error {
      db := config.DB
        id := c.Params("id")
        var book models.Book
      	result := db.First(&book, id)
	if result.Error != nil {
		// Handle the error, log it, and return a meaningful response
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
        return c.Render("download", fiber.Map{
            "Title":       book.Title,
            "Description": book.Description,
            "Image":       book.Image,
            "Link":        book.Link,
        })
    })
	app.Get("/display",controllers.DisplayAllBooks)
	
	app.Get("/admin",func(c *fiber.Ctx)error{
	  return c.Render("Create",fiber.Map{
	    "user":"x2",
	  })
	})
}
