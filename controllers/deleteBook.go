package controllers

import (
	"log"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"search/config"
	"search/models"
)

func DeleteMsgHandler(c *fiber.Ctx) error {
	// Extract the ID parameter from the URL path
	id := c.Params("id")

	// Convert the ID string to an integer
	msgID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("converting ID to integer:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	// Delete the message with the specified ID from the data store
	if err := config.DB.Delete(&models.Book{}, msgID).Error; err != nil {
		log.Println("deleting message:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not delete message"})
	}

	// Return a success response
	return c.JSON(fiber.Map{"message": "message deleted successfully"})
}
