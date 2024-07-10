package controllers

import (
	"github.com/gofiber/fiber/v2"
	"search/config"
	"search/models"
)

func GetAllBooks(c *fiber.Ctx) error {
	db := config.DB
	var books []models.Book
	result := db.Find(&books)
	if result.Error != nil {
		// Handle the error, log it, and return a meaningful response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(books)
}

func GetBookByID(c *fiber.Ctx) error {
	db := config.DB
	id := c.Params("id")
	var book models.Book
	result := db.First(&book, id)
	if result.Error != nil {
		// Handle the error, log it, and return a meaningful response
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}
