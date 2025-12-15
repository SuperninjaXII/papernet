package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"papernet/config"
	"papernet/model"

	"github.com/SuperninjaXII/goEpubTools/EpubTools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/minio"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"golang.org/x/crypto/bcrypt"
)

func CreateBook(c *fiber.Ctx) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	cartegories := c.FormValue("cartegories")
	date := c.FormValue("date")
	author := c.FormValue("author")
	coverImage, err := c.FormFile("image")
	if err != nil {
		log.Println("getting image error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid or missing image file",
		})
	}
	bookFile, err := c.FormFile("file")
	if err != nil {
		log.Println("getting file error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid or missing book file",
		})
	}
	bookFilePath := "./uploads/files/" + bookFile.Filename
	minioBookFilePath := "./files/" + bookFile.Filename
	MetaData := &epubtools.Package{
		MetaData: epubtools.MetaData{
			Title:       title,
			Description: description,
			Author:      author,
			Date:        date,
		},
	}
	epubtools.EditEpub(MetaData, bookFilePath, "./uploads/"+bookFile.Filename)
	err = c.SaveFile(bookFile, "./uploads/files/"+bookFile.Filename)
	if err != nil {
		log.Print("failed to save", err)
	}

	cover, err := coverImage.Open()
	if err != nil {
		log.Print(err)
	}
	defer cover.Close()
	//check file size
	coverBytes, err := io.ReadAll(cover)
	if err != nil {
		log.Print("failed to read image:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read image",
		})
	}

	file, err := bookFile.Open()
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	fileBuffer := make([]byte, 512)
	_, err = file.Read(fileBuffer)
	if err != nil {
		log.Print("error is ", err)
	}
	file.Seek(0, io.SeekStart)
	//minio to detect the io.reader
	bucketName := "uploads"
	minioClient := config.S3()
	// Ensure the bucket exists or create it
	if err := minioClient.CheckBucket(); err != nil {
		log.Println("Bucket check error:", err)

	}

	options := bimg.Options{
		Width:         1000,
		Height:        1600,
		Quality:       75,
		Compression:   9,
		Type:          bimg.WEBP,
		StripMetadata: true,
		Interpolator:  bimg.Bicubic,
	}
	processedImg, err := bimg.NewImage(coverBytes).Process(options)
	readImg := bytes.NewReader(processedImg)
	// Upload the cover image to MinIO
	coverImagePath := "images/" + title + ".webp"
	minio.ConfigDefault.PutObjectOptions.ContentType = "image/webp"

	_, err = minioClient.Conn().PutObject(
		c.Context(),
		bucketName,
		coverImagePath,
		readImg,
		int64(readImg.Len()),
		minio.ConfigDefault.PutObjectOptions,
	)
	if err != nil {
		log.Println("Cover image upload error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload cover image",
		})
	}
	if http.DetectContentType(fileBuffer) == "application/zip" {

		minio.ConfigDefault.PutObjectOptions.ContentType = "application/epub+zip"
	}
	// Upload the book file to MinIO
	_, err = minioClient.Conn().FPutObject(
		c.Context(),
		bucketName,
		minioBookFilePath,
		"./uploads/"+bookFile.Filename,
		minio.ConfigDefault.PutObjectOptions,
	)

	if err != nil {
		log.Println("Book file upload error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload book file",
		})
	}

	newImageName := title + ".webp"
	// Create a new book record with the uploaded file details
	newBook := model.Book{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		File:        bookFile.Filename,
		Image:       newImageName,
		Cartegories: cartegories,
		Author:      author,
	}

	// Save the book record in the database
	db := config.DB
	if err := db.Create(&newBook).Error; err != nil {
		log.Println("Database error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save book to database",
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book created successfully",
		"book":    newBook,
	})
}

// the bacrpt HashPassword( function
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(c *fiber.Ctx) error {
	var req model.User

	if err := c.BodyParser(&req); err != nil {
		log.Println("error with creating new user:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	if req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty password",
		})
	}
	newUUID := uuid.NewString()

	newPassword, err := HashPassword(req.Password)
	if err != nil || newPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Something went wrong with password ",
		})
	}

	newUser := model.User{
		ID:       newUUID,
		UserName: req.UserName,
		Email:    req.Email,
		Number:   req.Number,
		Password: newPassword,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		log.Println("error saving user to database:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"msg": "User created successfully",
	})
}
