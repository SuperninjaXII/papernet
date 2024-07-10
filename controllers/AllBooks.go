package controllers

import (
	"github.com/gofiber/fiber/v2"
	"search/config"
	"search/models"
	"fmt"
)

func DisplayAllBooks(c *fiber.Ctx) error {
    db := config.DB
    var books []models.Book
    result := db.Find(&books)
    if result.Error != nil {
        // Log the error
        fmt.Println("Error fetching books:", result.Error)
        // Return a meaningful response with status code 500
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
    }
    
    // Log the fetched books
    
    // Render the template with the fetched books
    return c.Render("book", fiber.Map{
        "Books": books,
    })
}



