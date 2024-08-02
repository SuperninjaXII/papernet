package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"papernet/config"
	"papernet/model"
)

// home page hoe page rendering
func Home(c *fiber.Ctx) error {
	return c.Render("layouts/mainLayout", fiber.Map{})
}

// search page rendering
func SearchPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

// books Page
func BooksPage(c *fiber.Ctx) error {

	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "query parameter is required"})
	}

	var books []model.Book
	err := config.DB.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&books).Error
	if err != nil {
		log.Println("error searching books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.Render("components/books", fiber.Map{
		"Result": books,
	})
}

// download page
func DownloadPage(c *fiber.Ctx) error {
	db := config.DB
	id := c.Params("id")
	var book model.Book
	result := db.First(&book, id)
	if result.Error != nil {
		// Handle the error, log it, and return a meaningful response
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.Render("components/download", fiber.Map{
		"book": book,
	})
}

//siggestions pushing to search page

func SearchResult(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "query parameter is required"})
	}

	var books []model.Book
	err := config.DB.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&books).Error
	if err != nil {
		log.Println("error searching books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.Render("components/searchResults", fiber.Map{
		"Result": books,
	})
}
