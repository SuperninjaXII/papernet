package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"papernet/config"
	"papernet/model"
)

func DisplayAllBooks(c *fiber.Ctx) error {
	db := config.DB
	var books []model.Book
	result := db.Find(&books)
	if result.Error != nil {
		// Log the error
		log.Println("Error fetching books:", result.Error)
		// Return a meaningful response with status code 500
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Log the fetched books

	// Render the template with the fetched books
	return c.Render("book", fiber.Map{
		"Books": books,
	})
}

func DisplayAllBooksJson(c *fiber.Ctx) error {
	db := config.DB
	var books []model.Book
	result := db.Find(&books)

	if result.Error != nil {
		// Log the error
		log.Println("Error fetching books:", result.Error)
		// Return a meaningful response with status code 500
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to fetch books",
		})
	}

	// Log the fetched books
	log.Println("Fetched books:", books)

	// Return the books as JSON response
	return c.JSON(fiber.Map{
		"books": books,
	})
}

func GetBookByID(c *fiber.Ctx) error {
	db := config.DB
	id := c.Params("id")
	var book model.Book
	result := db.First(&book, id)
	if result.Error != nil {
		// Handle the error, log it, and return a meaningful response
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}
