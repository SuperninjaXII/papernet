package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"papernet/config"
	"papernet/model"
)

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
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
