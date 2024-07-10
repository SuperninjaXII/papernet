package controllers

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "search/config"
    "search/models"
)

func UploadHandler(c *fiber.Ctx) error {
    var books []models.Book
    
    if err := c.BodyParser(&books); err != nil {
        log.Println("parsing error:", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "parsing request body"})
    }
    
    for _, book := range books {
        newBook := models.Book{
            Title:       book.Title,
            Description: book.Description,
            Link:        book.Link,
            Image: book.Image,
        }
        config.DB.Create(&newBook)
    }
    
    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "books added"})
}
