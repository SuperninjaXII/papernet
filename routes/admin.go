package routes

import (
	"papernet/controllers"

	"github.com/gofiber/fiber/v2"
)

// this package enambles funcs only admins should h haave
// to get authorization for tany  /admin functions use the jwt from the login routes
func Admin(app *fiber.App) {
	app.Post("/api/addBook", controllers.CreateBook)
	app.Post("/api/createuser", controllers.CreateUser)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/delete/:id", controllers.DeleteBookHandler)
}
