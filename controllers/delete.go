package controllers

import (
	"github.com/gofiber/fiber/v2"
	"papernet/config"
	"papernet/model"
)

func DeleteBookHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	// Validate the ID format (assuming it's a UUID)
	if len(id) != 36 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	// Delete the record
	result := config.DB.Delete(&model.Book{}, "ID = ?", id)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to delete",
		})
	}

	// Check if any record was deleted
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.JSON(fiber.Map{
		"msg": "Successful",
	})
}
