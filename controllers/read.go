package controllers

import (
	"log"
	"os"
	"papernet/config"
	"papernet/model"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Login function
func Login(c *fiber.Ctx) error {
	type LoginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var credentials LoginCredentials
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if credentials.Email == "" || credentials.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}
	var user model.User
	result := config.DB.First(&user, "email = ?", credentials.Email)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	claims := jwt.MapClaims{
		"userID": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 100).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SignedToken, err := token.SignedString([]byte(os.Getenv("jwt_key")))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": "jwt creation",
		})
	}
	return c.JSON(fiber.Map{
		"token":  SignedToken,
		"userID": user.ID,
	})
}

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
func DownloadBookFileByID(c *fiber.Ctx) error {
	db := config.DB
	id := c.Params("id")

	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	filePath := filepath.Join("uploads", "files", book.File)

	if err := c.Download(filePath); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}

	return nil
}
func DownloadCoverImageById(c *fiber.Ctx) error {
	db := config.DB
	id := c.Params("id")

	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	filePath := filepath.Join("uploads", "images", book.Image)

	if err := c.Download(filePath); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}

	return nil
}
