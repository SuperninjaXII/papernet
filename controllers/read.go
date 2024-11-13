package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"papernet/config"
	"papernet/model"
)

// Get all books
func GetAllBooks(c *fiber.Ctx) error {
	var books []model.Book
	if err := config.DB.Find(&books).Error; err != nil {
		log.Println("error fetching books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error fetching books"})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

// Search books by title or description

func SearchBooks(c *fiber.Ctx) error {
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
	return c.Status(fiber.StatusOK).JSON(books)
}
func GetBookByID(c *fiber.Ctx) error {
	db := config.DB
	id := c.Params("id")
	var book model.Book
	result := db.First(&book, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}
func GetBookByCartegory(c *fiber.Ctx) error {
    db := config.DB
    tag := c.Params("tag")
    var books []model.Book

    result := db.Where("cartegory1 = ? OR cartegory2 = ?", tag, tag).Find(&books)
    
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No books found in this category"})
    }

    return c.JSON(books)
}
