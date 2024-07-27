package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"papernet/config"
	"papernet/model"
)

func Create(c *fiber.Ctx) error {
	var books []model.Book

	if err := c.BodyParser(&books); err != nil {
		log.Println("parsing error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "parsing request body"})
	}

	for _, book := range books {
		newBook := model.Book{
			Title:       book.Title,
			Description: book.Description,
			Link:        book.Link,
			Image:       book.Image,
			Cartegory1:  book.Cartegory1,
			Cartegory2:  book.Cartegory2,
		}
		config.DB.Create(&newBook)
		log.Println("books created")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "books added"})
}

